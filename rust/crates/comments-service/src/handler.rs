use axum::{
    extract::{Path, State},
    response::{IntoResponse, Response},
};
use axum_htmx::HxRequest;
use color_eyre::eyre::WrapErr;
use tracing::instrument;

use crate::{
    model::{dto::CommentInfo, newtype::PositiveI64, template::BookComments},
    repo::{CommentRepo, CommentRepoImpl},
    AppState, ResponseError,
};

#[instrument(skip(state))]
pub async fn fetch_comments_for_book<A: AppState>(
    HxRequest(hx_request): HxRequest,
    Path(book_id): Path<PositiveI64>,
    State(state): State<A>,
) -> Result<Response, ResponseError>
where
    for<'a> &'a CommentRepoImpl<<A as AppState>::RepoInner>: CommentRepo,
{
    let book_id = book_id.get();
    let comments: Box<[CommentInfo]> = state
        .comment_repo()
        .fetch_comments_by_book_id(book_id)
        .await
        .wrap_err_with(|| format!("Fetching comments for the book {book_id}"))?;

    let comments: BookComments<'_> = comments.iter().collect();

    Ok(comments.into_response())
}
