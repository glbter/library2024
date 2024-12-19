use std::{
    borrow::Cow,
    cmp::Ordering,
    collections::{btree_map::Entry, BTreeMap},
};

use askama_axum::Template;
use itertools::Itertools;

use crate::model::{
    dto::CommentInfo,
    newtype::{CommentId, Uname},
};

#[derive(Debug, Template)]
#[cfg_attr(test, derive(PartialEq))]
#[template(path = "comment.html")]
pub struct BookComment<'a> {
    id: &'a CommentId,
    username: &'a Uname,
    text: &'a str,
    responses: Responses<'a>,
}

impl<'a> BookComment<'a> {
    pub const fn new(
        id: &'a CommentId,
        username: &'a Uname,
        text: &'a str,
        responses: Responses<'a>,
    ) -> Self {
        Self {
            id,
            username,
            text,
            responses,
        }
    }

    pub const fn fresh(id: &'a CommentId, username: &'a Uname, text: &'a str) -> Self {
        Self::new(
            id,
            username,
            text,
            Responses::Loaded(BookComments {
                response_to: None,
                items: BTreeMap::new(),
            }),
        )
    }
}

impl<'a> From<&'a CommentInfo> for BookComment<'a> {
    fn from(
        CommentInfo {
            id,
            username,
            text,
            has_responses,
            ..
        }: &'a CommentInfo,
    ) -> Self {
        Self {
            id,
            username,
            text,
            responses: if *has_responses {
                Responses::NotLoaded
            } else {
                Responses::Loaded(BookComments::default())
            },
        }
    }
}

#[derive(Debug, Default)]
#[cfg_attr(test, derive(PartialEq))]
pub enum Responses<'a> {
    /// There are responses, but they've not been loaded
    #[default]
    NotLoaded,
    /// All responses have been loaded, if any
    Loaded(BookComments<'a>),
}

impl<'a> Responses<'a> {
    #[must_use = "variant changes to `Responses::Loaded`"]
    pub fn as_loaded_mut(&mut self) -> &mut BookComments<'a> {
        match self {
            Self::NotLoaded => {
                *self = Self::Loaded(Default::default());
                let Self::Loaded(loaded) = self else {
                    unreachable!()
                };
                loaded
            }
            Self::Loaded(loaded) => loaded,
        }
    }
}

impl<'a> From<BookComments<'a>> for Responses<'a> {
    fn from(value: BookComments<'a>) -> Self {
        Self::Loaded(value)
    }
}

#[derive(Debug, Default, Template)]
#[cfg_attr(test, derive(PartialEq))]
#[template(path = "comments.html")]
pub struct BookComments<'a> {
    response_to: Option<&'a CommentId>,
    items: BTreeMap<&'a CommentId, BookComment<'a>>,
}

impl BookComments<'_> {
    pub fn list_id(&self) -> Cow<'static, str> {
        if let Some(id) = self.response_to {
            format!("responses-to-{id}").into()
        } else {
            "comments".into()
        }
    }
}

impl<'a> FromIterator<&'a CommentInfo> for BookComments<'a> {
    fn from_iter<T: IntoIterator<Item = &'a CommentInfo>>(iter: T) -> Self {
        let comments = iter
            .into_iter()
            .chunk_by(|c| c.response_to.as_ref())
            .into_iter()
            .sorted_unstable_by(|(a_response_to, _), (b_response_to, _)| {
                match (*a_response_to, *b_response_to) {
                    (None, None) => Ordering::Equal,
                    (None, _) => Ordering::Less,
                    (_, None) => Ordering::Greater,
                    (Some(a_response_to), Some(b_response_to)) => a_response_to.cmp(b_response_to),
                }
            })
            .fold(BTreeMap::new(), accumulate_comments);

        Self {
            items: comments,
            ..Self::default()
        }
    }
}

fn accumulate_comments<'a>(
    mut comments: BTreeMap<&'a CommentId, BookComment<'a>>,
    (response_to, responses): (Option<&'a CommentId>, impl Iterator<Item = &'a CommentInfo>),
) -> BTreeMap<&'a CommentId, BookComment<'a>> {
    match response_to {
        None => comments
            .extend(responses.map(|info @ CommentInfo { id, .. }| (id, BookComment::from(info)))),
        Some(response_to) => {
            if insert_responses(&mut comments, responses, response_to).is_some() {
                unreachable!("could not store responses")
            }
        }
    }

    comments
}

fn insert_responses<'a, I>(
    comments: &mut BTreeMap<&'a CommentId, BookComment<'a>>,
    mut responses: I,
    response_to: &'a CommentId,
) -> Option<I>
where
    I: Iterator<Item = &'a CommentInfo>,
{
    match comments.entry(response_to) {
        Entry::Vacant(_) => {
            for comment in comments.values_mut() {
                responses = insert_responses(
                    &mut comment.responses.as_loaded_mut().items,
                    responses,
                    response_to,
                )?;
            }
            Some(responses)
        }
        Entry::Occupied(mut response_to) => {
            response_to
                .get_mut()
                .responses
                .as_loaded_mut()
                .items
                .extend(
                    responses.map(|info @ CommentInfo { id, .. }| (id, BookComment::from(info))),
                );
            None
        }
    }
}

#[cfg(test)]
mod tests {
    use std::{iter, sync::LazyLock};

    use pretty_assertions::{assert_eq, assert_str_eq};
    use uuid::Uuid;

    use super::*;
    use crate::model::newtype::Username;

    static COMMENTS: LazyLock<(Vec<CommentId>, Vec<CommentInfo>)> = LazyLock::new(|| {
        let ids = Vec::from_iter(iter::from_fn(|| Some(CommentId::new(Uuid::now_v7()))).take(6));
        let comments = {
            vec![
                CommentInfo {
                    id: ids[0],
                    response_to: None,
                    username: Username::new("0"),
                    text: "".to_string(),
                    has_responses: false,
                },
                CommentInfo {
                    id: ids[1],
                    response_to: None,
                    username: Username::new("1"),
                    text: "".to_string(),
                    has_responses: true,
                },
                CommentInfo {
                    id: ids[2],
                    response_to: Some(ids[1]),
                    username: Username::new("2"),
                    text: "".to_string(),
                    has_responses: true,
                },
                CommentInfo {
                    id: ids[3],
                    response_to: Some(ids[2]),
                    username: Username::new("3"),
                    text: "".to_string(),
                    has_responses: false,
                },
                CommentInfo {
                    id: ids[4],
                    response_to: None,
                    username: Username::new("4"),
                    text: "".to_string(),
                    has_responses: false,
                },
                CommentInfo {
                    id: ids[5],
                    response_to: Some(ids[1]),
                    username: Username::new("5"),
                    text: "".to_string(),
                    has_responses: false,
                },
            ]
        };

        (ids, comments)
    });

    /// Collect complete list of comments into a tree
    #[test]
    fn collect_book_comments() {
        let ids = &*COMMENTS.0;
        let result: BookComments<'_> = COMMENTS.1.iter().collect();

        assert_eq!(
            result,
            BookComments {
                response_to: None,
                items: {
                    BTreeMap::from([
                        (
                            &ids[0],
                            BookComment {
                                id: &ids[0],
                                username: "0".into(),
                                text: "",
                                responses: BookComments::default().into(),
                            },
                        ),
                        (
                            &ids[1],
                            BookComment {
                                id: &ids[1],
                                username: "1".into(),
                                text: "",
                                responses: BookComments {
                                    response_to: Some(&ids[1]),
                                    items: BTreeMap::from([
                                        (
                                            &ids[2],
                                            BookComment {
                                                id: &ids[2],
                                                username: "2".into(),
                                                text: "",
                                                responses: BookComments {
                                                    response_to: Some(&ids[2]),
                                                    items: BTreeMap::from([(
                                                        &ids[3],
                                                        BookComment {
                                                            id: &ids[3],
                                                            username: "3".into(),
                                                            text: "",
                                                            responses: BookComments {
                                                                response_to: Some(&ids[3]),
                                                                ..BookComments::default()
                                                            }
                                                            .into(),
                                                        },
                                                    )]),
                                                }
                                                .into(),
                                            },
                                        ),
                                        (
                                            &ids[5],
                                            BookComment {
                                                id: &ids[5],
                                                username: "5".into(),
                                                text: "",
                                                responses: BookComments {
                                                    response_to: Some(&ids[5]),
                                                    ..BookComments::default()
                                                }
                                                .into(),
                                            },
                                        ),
                                    ]),
                                }
                                .into(),
                            },
                        ),
                        (
                            &ids[4],
                            BookComment {
                                id: &ids[4],
                                username: "4".into(),
                                text: "",
                                responses: BookComments {
                                    response_to: Some(&ids[4]),
                                    ..BookComments::default()
                                }
                                .into(),
                            },
                        ),
                    ])
                }
            }
        );
    }

    /// Some responses are not listed -> parent comments show them as [`NotLoaded`](Responses::NotLoaded)
    #[test]
    fn collect_book_comments_incomplete() {
        let ids = &*COMMENTS.0;
        let result: BookComments<'_> = COMMENTS
            .1
            .iter()
            .enumerate()
            .filter_map(|(i, c)| (i != 3).then_some(c)) // skip the 4th comment
            .collect();

        assert_eq!(
            result,
            BookComments {
                response_to: None,
                items: {
                    BTreeMap::from([
                        (
                            &ids[0],
                            BookComment {
                                id: &ids[0],
                                username: "0".into(),
                                text: "",
                                responses: BookComments::default().into(),
                            },
                        ),
                        (
                            &ids[1],
                            BookComment {
                                id: &ids[1],
                                username: "1".into(),
                                text: "",
                                responses: BookComments {
                                    response_to: Some(&ids[1]),
                                    items: BTreeMap::from([
                                        (
                                            &ids[2],
                                            BookComment {
                                                id: &ids[2],
                                                username: "2".into(),
                                                text: "",
                                                responses: Responses::NotLoaded,
                                            },
                                        ),
                                        (
                                            &ids[5],
                                            BookComment {
                                                id: &ids[5],
                                                username: "5".into(),
                                                text: "",
                                                responses: BookComments {
                                                    response_to: Some(&ids[5]),
                                                    ..BookComments::default()
                                                }
                                                .into(),
                                            },
                                        ),
                                    ]),
                                }
                                .into(),
                            },
                        ),
                        (
                            &ids[4],
                            BookComment {
                                id: &ids[4],
                                username: "4".into(),
                                text: "",
                                responses: BookComments {
                                    response_to: Some(&ids[4]),
                                    ..BookComments::default()
                                }
                                .into(),
                            },
                        ),
                    ])
                }
            }
        );
    }

    #[test]
    #[ignore = "unstable"]
    fn render_book_comments() {
        let comments: BookComments<'_> = COMMENTS.1.iter().collect();

        let result = comments.render().unwrap();

        assert_str_eq!(
            result,
            r"<ol>
  <li><div>
  <p>0</p>
  <p></p>
</div></li>
  <li><div>
  <p>1</p>
  <p></p>
  <ol>
  <li><div>
  <p>2</p>
  <p></p>
  <ol>
  <li><div>
  <p>3</p>
  <p></p>
</div></li>
</ol>
</div></li>
  <li><div>
  <p>5</p>
  <p></p>
</div></li>
</ol>
</div></li>
  <li><div>
  <p>4</p>
  <p></p>
</div></li>
</ol>"
        );
    }
}
