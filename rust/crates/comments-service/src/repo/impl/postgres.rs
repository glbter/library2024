use axum::async_trait;
use sqlx::{PgExecutor, PgPool};
use tracing::instrument;

use crate::{
    model::dto::CommentInfo,
    repo::{CommentRepo, CommentRepoImpl},
};

#[async_trait]
impl<'a> CommentRepo for &'a CommentRepoImpl<PgPool> {
    type Error = sqlx::Error;

    #[instrument(skip(self))]
    async fn fetch_comments_by_book_id(self, book_id: i64) -> sqlx::Result<Box<[CommentInfo]>> {
        fetch_comments_by_book_id(self.pool(), book_id).await
    }
}

#[instrument(skip(executor))]
async fn fetch_comments_by_book_id<'c>(
    executor: impl PgExecutor<'c>,
    book_id: i64,
) -> sqlx::Result<Box<[CommentInfo]>> {
    sqlx::query_as!(
            CommentInfo,
r#"SELECT c.id, c.response_to, c.text, concat_ws(' ', u.first_name, u.last_name) AS "username!"
FROM comments AS c
JOIN books AS b ON c.book_id = b.id
JOIN users AS u ON c.user_id = u.id
WHERE b.id = $1
ORDER BY c.id"#,
            book_id,
        )
        .fetch_all(executor)
        .await
        .map(Vec::into_boxed_slice)
}
