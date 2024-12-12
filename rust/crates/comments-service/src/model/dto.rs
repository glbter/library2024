use serde::Serialize;

use crate::model::newtype::CommentId;

#[derive(Debug, Serialize)]
pub struct CommentInfo {
    pub id: CommentId,
    pub response_to: Option<CommentId>,
    pub has_responses: bool,
    pub username: String,
    pub text: String,
}
