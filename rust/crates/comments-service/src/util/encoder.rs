use axum_extra::extract::cookie::Cookie;
use base64::{engine::GeneralPurpose, Engine};
use uuid::Uuid;

use crate::model::newtype::{SessionId, UserId};

static BASE64_ENGINE: GeneralPurpose = base64::engine::general_purpose::STANDARD;

pub struct SessionCookie {
    pub session_id: SessionId,
    pub user_id: UserId,
}

pub fn decode_session_cookie(
    cookie: &Cookie<'_>,
) -> Result<SessionCookie, DecodeSessionCookieError> {
    let decoded_value = BASE64_ENGINE.decode(cookie.value())?;
    Ok(SessionCookie {
        session_id: SessionId::new(Uuid::from_slice(&decoded_value[..16])?),
        user_id: UserId::new(i64::from_le_bytes(<[u8; 8]>::try_from(
            &decoded_value[16..32],
        )?)),
    })
}

#[derive(Debug, thiserror::Error)]
pub enum DecodeSessionCookieError {
    #[error(transparent)]
    Base64(#[from] base64::DecodeError),
    #[error(transparent)]
    Uuid(#[from] uuid::Error),
    #[error(transparent)]
    FromSlice(#[from] std::array::TryFromSliceError),
}
