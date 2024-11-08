mod r#impl;

use axum::async_trait;
use sqlx::{Database, Pool};

use crate::model::dto::CommentInfo;

#[derive(Debug, Clone)]
pub struct CommentRepoImpl<T>(T);

impl<D: Database> CommentRepoImpl<Pool<D>> {
    pub fn pool(&self) -> &Pool<D> {
        &self.0
    }
}

impl<T> CommentRepoImpl<T>
where
    for<'a> &'a Self: CommentRepo,
{
    pub fn new(inner: T) -> Self {
        Self(inner)
    }
}

#[async_trait]
pub trait CommentRepo {
    type Error: std::error::Error + Send + Sync + 'static;

    async fn fetch_comments_by_book_id(
        self,
        book_id: i64,
    ) -> Result<Box<[CommentInfo]>, Self::Error>;
}
