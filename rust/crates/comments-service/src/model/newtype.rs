use core::fmt;
use std::{hint, num::NonZero};

use serde::{
    de::{Error, Unexpected, Visitor},
    Deserialize, Deserializer,
};

#[derive(Debug, Clone, Copy, PartialEq, Eq, PartialOrd, Ord, Hash, sqlx::Type)]
#[sqlx(transparent)]
pub struct PositiveI64(NonZero<i64>);

impl PositiveI64 {
    pub const MIN: Self = unsafe { Self::new_unchecked(1) };
    pub const MAX: Self = unsafe { Self::new_unchecked(i64::MAX) };

    #[inline]
    pub const fn new(value: i64) -> Option<Self> {
        if value > 0 {
            Some(Self(unsafe { NonZero::new_unchecked(value) }))
        } else {
            None
        }
    }

    #[inline]
    pub const fn new_nonzero(value: NonZero<i64>) -> Option<Self> {
        if value.get() > 0 {
            Some(Self(value))
        } else {
            None
        }
    }

    /// # Safety
    /// Provided value must be greater than 0.
    #[inline]
    pub const unsafe fn new_unchecked(value: i64) -> Self {
        if let Some(positive) = Self::new(value) {
            positive
        } else {
            hint::unreachable_unchecked()
        }
    }

    #[inline]
    pub fn get(self) -> i64 {
        self.0.get()
    }
}

impl TryFrom<i64> for PositiveI64 {
    type Error = NonPositive;

    fn try_from(value: i64) -> Result<Self, Self::Error> {
        PositiveI64::new(value).ok_or(NonPositive)
    }
}

impl TryFrom<NonZero<i64>> for PositiveI64 {
    type Error = NonPositive;

    fn try_from(value: NonZero<i64>) -> Result<Self, Self::Error> {
        PositiveI64::new_nonzero(value).ok_or(NonPositive)
    }
}

#[derive(Debug, thiserror::Error)]
#[error("number provided is not positive")]
pub struct NonPositive;

impl<'de> Deserialize<'de> for PositiveI64 {
    fn deserialize<D>(deserializer: D) -> Result<Self, D::Error>
    where
        D: Deserializer<'de>,
    {
        struct PositiveVisitor;

        impl<'de> Visitor<'de> for PositiveVisitor {
            type Value = PositiveI64;

            fn expecting(&self, f: &mut fmt::Formatter) -> fmt::Result {
                f.write_str(concat!("a positive ", stringify!(i64)))
            }

            fn visit_i64<E>(self, v: i64) -> Result<Self::Value, E>
            where
                E: Error,
            {
                match Self::Value::new(v) {
                    Some(positive) => Ok(positive),
                    None => Err(Error::invalid_value(Unexpected::Signed(v), &self)),
                }
            }

            fn visit_u64<E>(self, v: u64) -> Result<Self::Value, E>
            where
                E: Error,
            {
                if v <= i64::MAX as u64 {
                    if let Some(nonzero) = Self::Value::new(v as i64) {
                        return Ok(nonzero);
                    }
                }
                Err(Error::invalid_value(Unexpected::Unsigned(v), &self))
            }
        }

        deserializer.deserialize_i64(PositiveVisitor)
    }
}
