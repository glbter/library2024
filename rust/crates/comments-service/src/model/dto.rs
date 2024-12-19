use serde::Serialize;

use crate::model::newtype::{CommentId, Username};

#[derive(Debug, Serialize)]
pub struct CommentInfo {
    pub id: CommentId,
    pub response_to: Option<CommentId>,
    pub has_responses: bool,
    pub username: Username,
    pub text: String,
}
