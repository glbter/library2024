use core::fmt::{self, Formatter};

use uuid::Uuid;

#[derive(
    Debug,
    Clone,
    Copy,
    PartialEq,
    Eq,
    PartialOrd,
    Ord,
    Hash,
    sqlx::Type,
    serde::Serialize,
    serde::Deserialize,
)]
#[sqlx(transparent)]
#[serde(transparent)]
#[repr(transparent)]
pub struct BookId(i64);

impl BookId {
    pub const fn new(value: i64) -> Self {
        Self(value)
    }

    pub const fn get(self) -> i64 {
        self.0
    }
}

impl fmt::Display for BookId {
    fn fmt(&self, f: &mut Formatter<'_>) -> fmt::Result {
        fmt::Display::fmt(&self.0, f)
    }
}

#[derive(
    Debug,
    Clone,
    Copy,
    PartialEq,
    Eq,
    PartialOrd,
    Ord,
    Hash,
    sqlx::Type,
    serde::Serialize,
    serde::Deserialize,
)]
#[sqlx(transparent)]
#[serde(transparent)]
#[repr(transparent)]
pub struct CommentId(Uuid);

impl CommentId {
    pub const fn new(value: Uuid) -> Self {
        Self(value)
    }

    pub const fn get(self) -> Uuid {
        self.0
    }
}

impl fmt::Display for CommentId {
    fn fmt(&self, f: &mut Formatter<'_>) -> fmt::Result {
        fmt::Display::fmt(&self.0, f)
    }
}

#[derive(
    Debug,
    Clone,
    Copy,
    PartialEq,
    Eq,
    PartialOrd,
    Ord,
    Hash,
    sqlx::Type,
    serde::Serialize,
    serde::Deserialize,
)]
#[sqlx(transparent)]
#[serde(transparent)]
#[repr(transparent)]
pub struct UserId(i64);

impl UserId {
    pub const fn new(value: i64) -> Self {
        Self(value)
    }

    pub const fn get(self) -> i64 {
        self.0
    }
}

impl fmt::Display for UserId {
    fn fmt(&self, f: &mut Formatter<'_>) -> fmt::Result {
        fmt::Display::fmt(&self.0, f)
    }
}

#[derive(
    Debug,
    Clone,
    Copy,
    PartialEq,
    Eq,
    PartialOrd,
    Ord,
    Hash,
    sqlx::Type,
    serde::Serialize,
    serde::Deserialize,
)]
#[sqlx(transparent)]
#[serde(transparent)]
#[repr(transparent)]
pub struct SessionId(Uuid);

impl SessionId {
    pub const fn new(value: Uuid) -> Self {
        Self(value)
    }

    pub const fn get(self) -> Uuid {
        self.0
    }
}

impl fmt::Display for SessionId {
    fn fmt(&self, f: &mut Formatter<'_>) -> fmt::Result {
        fmt::Display::fmt(&self.0, f)
    }
}
