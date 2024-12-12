use axum::{async_trait, extract::FromRequestParts, http::request::Parts, RequestPartsExt};
use axum_extra::extract::CookieJar;
use color_eyre::eyre::WrapErr;

use crate::{
    model::newtype::UserId,
    repo::{SessionRepo, SessionRepoImpl},
    util::{encoder, encoder::SessionCookie},
    AppState, ResponseError,
};

#[derive(Debug, Clone, Copy, PartialEq, Eq, PartialOrd, Ord)]
pub struct User(UserId);

impl User {
    pub const fn id(self) -> UserId {
        self.0
    }
}

#[async_trait]
impl<S: AppState> FromRequestParts<S> for User
where
    for<'a> &'a SessionRepoImpl<<S as AppState>::RepoInner>: SessionRepo,
{
    type Rejection = ResponseError;

    async fn from_request_parts(parts: &mut Parts, state: &S) -> Result<Self, Self::Rejection> {
        let cookies = parts.extract::<CookieJar>().await?;
        let session_cookie =
            cookies
                .get(state.session_cookie_name())
                .ok_or(ResponseError::Unauthorized {
                    reason: "session cookie absent".into(),
                })?;

        let SessionCookie {
            session_id,
            user_id,
        } = encoder::decode_session_cookie(session_cookie).map_err(|err| {
            ResponseError::Unauthorized {
                reason: err.to_string().into(),
            }
        })?;

        let valid = state
            .session_repo()
            .validate(session_id, user_id)
            .await
            .wrap_err("Failed to validate session repo")?;

        if valid {
            Ok(User(user_id))
        } else {
            Err(ResponseError::Unauthorized {
                reason: "invalid session token".into(),
            })
        }
    }
}
