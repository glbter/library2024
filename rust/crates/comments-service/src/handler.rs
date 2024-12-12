use axum::{
    body::Body,
    extract::{Path, Query, State},
    http::header,
    response::{IntoResponse, Response},
    Form,
};
use color_eyre::eyre::WrapErr;
use serde::Deserialize;
use tracing::instrument;

use crate::{
    extractor::User,
    model::{
        dto::CommentInfo,
        newtype::{BookId, CommentId},
        template::BookComments,
    },
    repo::{CommentRepo, CommentRepoImpl},
    AppState, ResponseError,
};

#[derive(Debug, Deserialize)]
#[serde(default)]
pub struct FetchCommentsQuery {
    toplevel: bool,
}

impl Default for FetchCommentsQuery {
    fn default() -> Self {
        FetchCommentsQuery { toplevel: true }
    }
}

#[instrument(skip(state))]
pub async fn get_comments_for_book<S: AppState>(
    State(state): State<S>,
    Path(book_id): Path<BookId>,
    Query(FetchCommentsQuery { toplevel }): Query<FetchCommentsQuery>,
) -> Result<Response, ResponseError>
where
    for<'a> &'a CommentRepoImpl<<S as AppState>::RepoInner>: CommentRepo,
{
    let repo = state.comment_repo();
    let comments: Box<[CommentInfo]> = if toplevel {
        repo.select_toplevel_by_book_id(book_id)
            .await
            .wrap_err_with(|| format!("Fetching toplevel comments for the book {book_id}"))?
    } else {
        repo.select_by_book_id(book_id)
            .await
            .wrap_err_with(|| format!("Fetching comments for the book {book_id}"))?
    };

    let comments: BookComments<'_> = comments.iter().collect();

    Ok(comments.into_response())
}

#[instrument(skip(state))]
pub async fn get_comment_responses<S: AppState>(
    State(state): State<S>,
    Path(comment_id): Path<CommentId>,
) -> Result<Response, ResponseError>
where
    for<'a> &'a CommentRepoImpl<<S as AppState>::RepoInner>: CommentRepo,
{
    let comments: Box<[CommentInfo]> = state
        .comment_repo()
        .select_responses_by_id(comment_id)
        .await
        .wrap_err_with(|| format!("Fetching responses for comment {comment_id}"))?;

    let comments: BookComments<'_> = comments.iter().collect();

    Ok(comments.into_response())
}

#[derive(Debug, Deserialize)]
pub struct PostCommentOnBookForm {
    text: String,
}

#[instrument(skip(state))]
pub async fn post_comment_on_book<S: AppState>(
    State(state): State<S>,
    user: User,
    Path(book_id): Path<BookId>,
    Form(PostCommentOnBookForm { text }): Form<PostCommentOnBookForm>,
) -> Result<Response, ResponseError>
where
    for<'a> &'a CommentRepoImpl<<S as AppState>::RepoInner>: CommentRepo,
{
    let user_id = user.id();
    let id = state
        .comment_repo()
        .insert_by_book_id(book_id, user_id, text)
        .await
        .wrap_err_with(|| format!("inserting comment for a book {book_id} by user {user_id}"))?;
    Ok(Response::builder()
        .header(header::LOCATION, format!("/api/comments/{id}"))
        .body(Body::empty())
        .unwrap())
}

#[instrument(skip(state))]
pub async fn post_comment_response<S: AppState>(
    State(state): State<S>,
    user: User,
    Path(comment_id): Path<CommentId>,
    Form(PostCommentOnBookForm { text }): Form<PostCommentOnBookForm>,
) -> Result<Response, ResponseError>
where
    for<'a> &'a CommentRepoImpl<<S as AppState>::RepoInner>: CommentRepo,
{
    let user_id = user.id();
    let id = state
        .comment_repo()
        .insert_by_response_to_id(comment_id, user_id, text)
        .await
        .wrap_err_with(|| {
            format!("inserting response to a comment {comment_id} by user {user_id}")
        })?;
    Ok(Response::builder()
        .header(header::LOCATION, format!("/api/comments/{id}"))
        .body(Body::empty())
        .unwrap())
}
