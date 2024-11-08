use serde::Serialize;
use uuid::Uuid;

#[derive(Debug, Serialize)]
pub struct CommentInfo {
    pub id: Uuid,
    pub response_to: Option<Uuid>,
    pub username: String,
    pub text: String,
}
