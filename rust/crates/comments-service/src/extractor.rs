use axum::{async_trait, extract::FromRequestParts, http::request::Parts, RequestPartsExt};
use axum_extra::extract::CookieJar;
use color_eyre::eyre::WrapErr;
use tracing::instrument;

use crate::{
    model::newtype::{Uname, UserId, Username},
    repo::{SessionRepo, SessionRepoImpl},
    util::{encoder, encoder::SessionCookie},
    AppState, ResponseError,
};

#[derive(Debug)]
pub struct User {
    id: UserId,
    name: Username,
}

impl User {
    pub const fn id(&self) -> UserId {
        self.id
    }

    pub fn name(&self) -> &Uname {
        &self.name
    }
}

#[async_trait]
impl<S: AppState> FromRequestParts<S> for User
where
    for<'a> &'a SessionRepoImpl<<S as AppState>::RepoInner>: SessionRepo,
{
    type Rejection = ResponseError;

    #[instrument(skip_all)]
    async fn from_request_parts(parts: &mut Parts, state: &S) -> Result<Self, Self::Rejection> {
        let cookies = parts.extract::<CookieJar>().await?;
        let session_cookie =
            cookies
                .get(state.session_cookie_name())
                .ok_or(ResponseError::Unauthorized {
                    reason: "session cookie absent".into(),
                })?;

        tracing::debug!("Obtained session cookie");

        let session_cookie @ SessionCookie {
            session_id,
            user_id,
        } = encoder::decode_session_cookie(session_cookie).map_err(|err| {
            ResponseError::Unauthorized {
                reason: err.to_string().into(),
            }
        })?;

        tracing::debug!("{session_cookie:?}");

        let username = state
            .session_repo()
            .select_username(session_id, user_id)
            .await
            .wrap_err("Failed to validate session repo")?;

        tracing::debug!("Username: {username:?}");

        if let Some(username) = username {
            Ok(User {
                id: user_id,
                name: username,
            })
        } else {
            Err(ResponseError::Unauthorized {
                reason: "invalid session token".into(),
            })
        }
    }
}
