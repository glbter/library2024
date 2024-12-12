use axum::async_trait;
use sqlx::{PgExecutor, PgPool};
use tracing::instrument;
use uuid::Uuid;

use crate::{
    model::{
        dto::CommentInfo,
        newtype::{BookId, CommentId, SessionId, UserId},
    },
    repo::{CommentRepo, CommentRepoImpl, SessionRepo, SessionRepoImpl},
};

#[async_trait]
impl CommentRepo for &CommentRepoImpl<PgPool> {
    type Error = sqlx::Error;

    #[instrument(skip(self))]
    async fn select_by_book_id(self, book_id: BookId) -> sqlx::Result<Box<[CommentInfo]>> {
        select_comments_by_book_id(self.pool(), book_id).await
    }

    #[instrument(skip(self))]
    async fn select_toplevel_by_book_id(self, book_id: BookId) -> sqlx::Result<Box<[CommentInfo]>> {
        select_toplevel_comments_by_book_id(self.pool(), book_id).await
    }

    #[instrument(skip(self))]
    async fn select_responses_by_id(
        self,
        comment_id: CommentId,
    ) -> Result<Box<[CommentInfo]>, Self::Error> {
        select_comment_responses_by_id(self.pool(), comment_id).await
    }

    #[instrument(skip(self))]
    async fn insert_by_book_id(
        self,
        book_id: BookId,
        user_id: UserId,
        text: String,
    ) -> Result<CommentId, Self::Error> {
        insert_comment_by_book_id(self.pool(), book_id, user_id, text).await
    }

    #[instrument(skip(self))]
    async fn insert_by_response_to_id(
        self,
        response_to: CommentId,
        user_id: UserId,
        text: String,
    ) -> Result<CommentId, Self::Error> {
        insert_comment_by_response_to_id(self.pool(), response_to, user_id, text).await
    }
}

#[instrument(skip(executor))]
async fn select_comments_by_book_id<'c>(
    executor: impl PgExecutor<'c>,
    book_id: BookId,
) -> sqlx::Result<Box<[CommentInfo]>> {
    sqlx::query_as!(
        CommentInfo,
        r#"SELECT
    c.id as "id: CommentId",
    c.response_to as "response_to: CommentId",
    c.text,
    concat_ws(' ', u.first_name, u.last_name) AS "username!",
    EXISTS (SELECT 1 FROM comments AS c1 WHERE c.id = c1.response_to) AS "has_responses!"
FROM comments AS c
JOIN books AS b ON c.book_id = b.id
JOIN users AS u ON c.user_id = u.id
WHERE b.id = $1
ORDER BY c.id"#,
        book_id as BookId,
    )
    .fetch_all(executor)
    .await
    .map(Vec::into_boxed_slice)
}

#[instrument(skip(executor))]
async fn select_toplevel_comments_by_book_id<'c>(
    executor: impl PgExecutor<'c>,
    book_id: BookId,
) -> sqlx::Result<Box<[CommentInfo]>> {
    sqlx::query_as!(
        CommentInfo,
        r#"SELECT
    c.id as "id: CommentId",
    NULL as "response_to: CommentId",
    c.text,
    concat_ws(' ', u.first_name, u.last_name) AS "username!",
    EXISTS (SELECT 1 FROM comments AS c1 WHERE c.id = c1.response_to) AS "has_responses!"
FROM comments AS c
JOIN books AS b ON c.book_id = b.id
JOIN users AS u ON c.user_id = u.id
WHERE b.id = $1 AND c.response_to = NULL
ORDER BY c.id"#,
        book_id as BookId,
    )
    .fetch_all(executor)
    .await
    .map(Vec::into_boxed_slice)
}

#[instrument(skip(executor))]
async fn select_comment_responses_by_id<'c>(
    executor: impl PgExecutor<'c>,
    comment_id: CommentId,
) -> sqlx::Result<Box<[CommentInfo]>> {
    sqlx::query_as!(
        CommentInfo,
        r#"SELECT
    c.id as "id: CommentId",
    c.response_to as "response_to: CommentId",
    c.text,
    concat_ws(' ', u.first_name, u.last_name) AS "username!",
    EXISTS (SELECT 1 FROM comments AS c1 WHERE c.id = c1.response_to) AS "has_responses!"
FROM comments AS c
JOIN users AS u ON c.user_id = u.id
WHERE c.response_to = $1
ORDER BY c.id"#,
        comment_id as CommentId,
    )
    .fetch_all(executor)
    .await
    .map(Vec::into_boxed_slice)
}

async fn insert_comment_by_book_id<'c>(
    executor: impl PgExecutor<'c>,
    book_id: BookId,
    user_id: UserId,
    text: String,
) -> sqlx::Result<CommentId> {
    let id = CommentId::new(Uuid::now_v7());
    sqlx::query!(
        r#"INSERT INTO comments (id, book_id, user_id, text) VALUES ($1, $2, $3, $4)"#,
        id as CommentId,
        book_id as BookId,
        user_id as UserId,
        text
    )
    .execute(executor)
    .await
    .map(|_| id)
}

async fn insert_comment_by_response_to_id<'c>(
    executor: impl PgExecutor<'c>,
    comment_id: CommentId,
    user_id: UserId,
    text: String,
) -> sqlx::Result<CommentId> {
    let id = CommentId::new(Uuid::now_v7());
    sqlx::query!(
        r#"WITH b AS (
    SELECT b.id
    FROM books AS b
    JOIN comments AS c ON c.book_id = b.id
    WHERE c.id = $2)
INSERT INTO comments (id, book_id, response_to, user_id, text)
VALUES ($1, (SELECT b.id FROM b LIMIT 1), $2, $3, $4)"#,
        id as CommentId,
        comment_id as CommentId,
        user_id as UserId,
        text
    )
    .execute(executor)
    .await
    .map(|_| id)
}

#[async_trait]
impl SessionRepo for &SessionRepoImpl<PgPool> {
    type Error = sqlx::Error;

    async fn validate(self, session_id: SessionId, user_id: UserId) -> Result<bool, Self::Error> {
        validate_session_user(self.pool(), session_id, user_id).await
    }
}

#[instrument(skip(executor))]
async fn validate_session_user<'c>(
    executor: impl PgExecutor<'c>,
    session_id: SessionId,
    user_id: UserId,
) -> sqlx::Result<bool> {
    sqlx::query!(
        r#"SELECT
    (s.user_id = $2) as "valid!"
FROM sessions AS s
WHERE s.id = $1"#,
        session_id as SessionId,
        user_id as UserId,
    )
    .fetch_one(executor)
    .await
    .map(|r| r.valid)
}
