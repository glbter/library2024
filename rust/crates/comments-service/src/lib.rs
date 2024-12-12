#![deny(rust_2018_idioms)]

use core::fmt;
use std::{borrow::Cow, convert::Infallible};

use askama_axum::Response;
use axum::{
    body::Body,
    http::StatusCode,
    response::IntoResponse,
    routing::{get, post},
    Router,
};
use sqlx::PgPool;
use tower_http::trace::TraceLayer;

use crate::repo::{CommentRepo, CommentRepoImpl, SessionRepo, SessionRepoImpl};

pub mod extractor;
pub mod handler;
pub mod model;
pub mod repo;
pub(crate) mod util;

#[derive(Debug, thiserror::Error)]
pub enum ResponseError {
    #[error("Not found {subject:?}")]
    NotFound { subject: &'static str },
    #[error("Unauthorized access: {reason}")]
    Unauthorized { reason: Cow<'static, str> },
    #[error("Internal Server Error")]
    Other(
        #[from]
        #[source]
        color_eyre::Report,
    ),
}

impl From<Infallible> for ResponseError {
    fn from(value: Infallible) -> Self {
        match value {}
    }
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
            ResponseError::Unauthorized { .. } => StatusCode::UNAUTHORIZED,
            ResponseError::Other(_) => StatusCode::INTERNAL_SERVER_ERROR,
        }
    }

    pub fn detailed_display_fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        match self {
            ResponseError::NotFound { .. } | ResponseError::Unauthorized { .. } => {
                fmt::Display::fmt(self, f)
            }
            ResponseError::Other(report) => fmt::Display::fmt(report, f),
        }
    }
}

pub trait AppState
where
    Self: Clone + Sync + Send + 'static,
{
    type RepoInner;

    fn comment_repo(&self) -> &CommentRepoImpl<Self::RepoInner>;

    fn session_repo(&self) -> &SessionRepoImpl<Self::RepoInner>;

    fn session_cookie_name(&self) -> &str;
}

pub fn build<S: AppState>(state: S) -> Router
where
    for<'a> &'a CommentRepoImpl<<S as AppState>::RepoInner>: CommentRepo,
    for<'a> &'a SessionRepoImpl<<S as AppState>::RepoInner>: SessionRepo,
{
    Router::new()
        .route(
            "/api/books/:book_id/comments",
            get(handler::get_comments_for_book::<S>),
        )
        .route(
            "/api/books/:book_id/comments",
            post(handler::post_comment_on_book::<S>),
        )
        .route(
            "/api/comments/:comment_id/responses",
            get(handler::get_comment_responses::<S>),
        )
        .route(
            "/api/comments/:comment_id/responses",
            post(handler::post_comment_response::<S>),
        )
        .layer(TraceLayer::new_for_http())
        .with_state(state)
}

#[derive(Clone)]
pub struct ProdAppState {
    comment_repo: CommentRepoImpl<PgPool>,
    session_repo: SessionRepoImpl<PgPool>,
    session_cookie_name: String,
}

impl ProdAppState {
    pub fn new(
        comment_repo: CommentRepoImpl<PgPool>,
        session_repo: SessionRepoImpl<PgPool>,
        session_cookie_name: String,
    ) -> Self {
        Self {
            comment_repo,
            session_repo,
            session_cookie_name,
        }
    }
}

impl AppState for ProdAppState {
    type RepoInner = PgPool;

    fn comment_repo(&self) -> &CommentRepoImpl<Self::RepoInner> {
        &self.comment_repo
    }

    fn session_repo(&self) -> &SessionRepoImpl<Self::RepoInner> {
        &self.session_repo
    }

    fn session_cookie_name(&self) -> &str {
        &self.session_cookie_name
    }
}
