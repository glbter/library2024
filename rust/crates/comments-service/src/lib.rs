extern crate core;

use core::fmt;

use askama_axum::Response;
use axum::{body::Body, http::StatusCode, response::IntoResponse, routing::get, Router};
use sqlx::{Database, PgPool};

use crate::repo::{CommentRepo, CommentRepoImpl};

pub mod handler;
pub mod model;
pub mod repo;

#[derive(Debug, thiserror::Error)]
pub enum ResponseError {
    #[error("{subject} not found")]
    NotFound { subject: &'static str },
    #[error("Internal Server Error")]
    Other(
        #[from]
        #[source]
        color_eyre::Report,
    ),
}

impl IntoResponse for ResponseError {
    fn into_response(self) -> Response {
        Response::builder()
            .status(self.status())
            .body(Body::new(self.to_string()))
            .unwrap()
    }
}

impl ResponseError {
    pub const fn status(&self) -> StatusCode {
        match self {
            ResponseError::NotFound { .. } => StatusCode::NOT_FOUND,
            ResponseError::Other(_) => StatusCode::INTERNAL_SERVER_ERROR,
        }
    }

    pub fn detailed_display_fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        match self {
            ResponseError::NotFound { .. } => fmt::Display::fmt(self, f),
            ResponseError::Other(report) => fmt::Display::fmt(report, f),
        }
    }
}

pub trait AppState
where
    Self: Clone + Sync + Send + 'static,
    for<'a> &'a CommentRepoImpl<Self::RepoInner>: CommentRepo,
{
    type RepoInner;

    fn comment_repo(&self) -> &CommentRepoImpl<Self::RepoInner>;
}

pub fn build<A: AppState>(state: A) -> Router
where
    for<'a> &'a CommentRepoImpl<<A as AppState>::RepoInner>: CommentRepo,
{
    Router::new()
        .route(
            "/books/:book_id/comments",
            get(handler::fetch_comments_for_book::<A>),
        )
        .with_state(state)
}

#[derive(Clone)]
pub struct ProdAppState {
    comment_repo: CommentRepoImpl<PgPool>,
}

impl ProdAppState {
    pub fn new(comment_repo: CommentRepoImpl<PgPool>) -> Self {
        Self { comment_repo }
    }
}

impl AppState for ProdAppState {
    type RepoInner = PgPool;

    fn comment_repo(&self) -> &CommentRepoImpl<Self::RepoInner> {
        &self.comment_repo
    }
}
