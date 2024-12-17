use core::fmt::{self, Formatter};
use std::ops::Deref;

use serde::{Deserialize, Deserializer};
use sqlx::{
    encode::IsNull,
    error::BoxDynError,
    postgres::{PgHasArrayType, PgTypeInfo},
    Database, Postgres,
};
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
pub struct Username(Box<Uname>);

impl Username {
    #[inline(always)]
    pub fn new(value: impl Into<Box<Uname>>) -> Self {
        Self(value.into())
    }

    pub const fn get(&self) -> &Uname {
        &self.0
    }

    pub fn into_inner(self) -> Box<Uname> {
        self.0
    }
}

impl fmt::Display for Username {
    fn fmt(&self, f: &mut Formatter<'_>) -> fmt::Result {
        fmt::Display::fmt(&self.0, f)
    }
}

impl From<String> for Username {
    #[inline(always)]
    fn from(s: String) -> Self {
        Self::new(s)
    }
}

#[derive(Debug, PartialEq, Eq, PartialOrd, Ord, Hash, serde::Serialize)]
#[repr(transparent)]
pub struct Uname(str);

impl Uname {
    pub const fn get(&self) -> &str {
        &self.0
    }
}

impl fmt::Display for Uname {
    fn fmt(&self, f: &mut Formatter<'_>) -> fmt::Result {
        fmt::Display::fmt(&self.0, f)
    }
}

impl Deref for Username {
    type Target = Uname;

    fn deref(&self) -> &Self::Target {
        unsafe { &*(self.0.deref() as *const _ as *const Self::Target) }
    }
}

impl From<&str> for &Uname {
    fn from(value: &str) -> Self {
        unsafe { &*(value as *const _ as *const Uname) }
    }
}

impl From<Box<str>> for Box<Uname> {
    fn from(value: Box<str>) -> Self {
        let ptr = Box::into_raw(value);
        unsafe { Box::from_raw(ptr as *mut Uname) }
    }
}

impl From<&str> for Box<Uname> {
    fn from(value: &str) -> Self {
        value.to_string().into_boxed_str().into()
    }
}

impl From<Box<Uname>> for Box<str> {
    fn from(value: Box<Uname>) -> Self {
        let ptr = Box::into_raw(value);
        unsafe { Box::from_raw(ptr as *mut str) }
    }
}

impl<'q> sqlx::Encode<'q, Postgres> for Box<Uname> {
    #[inline(always)]
    fn encode_by_ref(
        &self,
        buf: &mut <Postgres as Database>::ArgumentBuffer<'q>,
    ) -> Result<IsNull, BoxDynError> {
        let box_str = unsafe { std::mem::transmute(self) };
        <Box<str> as sqlx::Encode<'q, Postgres>>::encode_by_ref(box_str, buf)
    }
}

impl<'r> sqlx::Decode<'r, Postgres> for Box<Uname> {
    #[inline]
    fn decode(value: <Postgres as Database>::ValueRef<'r>) -> Result<Self, BoxDynError> {
        <Box<str> as sqlx::Decode<Postgres>>::decode(value).map(Box::from)
    }
}

impl PgHasArrayType for Box<Uname> {
    fn array_type_info() -> PgTypeInfo {
        Box::<str>::array_type_info()
    }
}

impl<'de> Deserialize<'de> for Box<Uname> {
    #[inline]
    fn deserialize<D>(deserializer: D) -> Result<Self, D::Error>
    where
        D: Deserializer<'de>,
    {
        Box::<str>::deserialize(deserializer).map(Box::from)
    }
}

impl Clone for Box<Uname> {
    fn clone(&self) -> Self {
        Box::<str>::clone(unsafe { std::mem::transmute(self) }).into()
    }
}

impl From<String> for Box<Uname> {
    fn from(value: String) -> Self {
        value.into_boxed_str().into()
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
