mod r#impl;

use axum::async_trait;
use sqlx::{Database, Pool};

use crate::model::{
    dto::CommentInfo,
    newtype::{BookId, CommentId, SessionId, UserId},
};

#[derive(Debug, Clone)]
pub struct CommentRepoImpl<T>(T);

impl<T> CommentRepoImpl<T>
where
    for<'a> &'a Self: CommentRepo,
{
    pub fn new(inner: T) -> Self {
        Self(inner)
    }
}

impl<D: Database> CommentRepoImpl<Pool<D>> {
    pub fn pool(&self) -> &Pool<D> {
        &self.0
    }
}

#[derive(Debug, Clone)]
pub struct SessionRepoImpl<T>(T);

impl<T> SessionRepoImpl<T>
where
    for<'a> &'a Self: SessionRepo,
{
    pub fn new(inner: T) -> Self {
        Self(inner)
    }
}

impl<D: Database> SessionRepoImpl<Pool<D>> {
    pub fn pool(&self) -> &Pool<D> {
        &self.0
    }
}

#[async_trait]
pub trait CommentRepo {
    type Error: std::error::Error + Send + Sync + 'static;

    async fn select_by_book_id(self, book_id: BookId) -> Result<Box<[CommentInfo]>, Self::Error>;

    async fn select_toplevel_by_book_id(
        self,
        book_id: BookId,
    ) -> Result<Box<[CommentInfo]>, Self::Error>;

    async fn select_responses_by_id(
        self,
        comment_id: CommentId,
    ) -> Result<Box<[CommentInfo]>, Self::Error>;

    async fn insert_by_book_id(
        self,
        book_id: BookId,
        user_id: UserId,
        text: &str,
    ) -> Result<CommentId, Self::Error>;

    async fn insert_by_response_to_id(
        self,
        response_to: CommentId,
        user_id: UserId,
        text: &str,
    ) -> Result<CommentId, Self::Error>;
}

#[async_trait]
pub trait SessionRepo {
    type Error: std::error::Error + Send + Sync + 'static;

    async fn select_username(
        self,
        session_id: SessionId,
        user_id: UserId,
    ) -> Result<Option<Box<str>>, Self::Error>;
}
