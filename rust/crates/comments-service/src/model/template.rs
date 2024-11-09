use askama_axum::Template;
use itertools::Itertools;
use std::ops::Deref;
use std::{
    cmp::Ordering,
    collections::{btree_map::Entry, BTreeMap},
};
use uuid::Uuid;

use crate::model::dto::CommentInfo;

#[derive(Debug, Template)]
#[cfg_attr(test, derive(PartialEq))]
#[template(path = "comment.html")]
pub struct BookComment<'a> {
    username: &'a str,
    text: &'a str,
    responses: BookComments<'a>,
}

#[derive(Debug, Default, Template)]
#[cfg_attr(test, derive(PartialEq))]
#[template(path = "comments.html")]
#[repr(transparent)]
pub struct BookComments<'a>(BTreeMap<Uuid, BookComment<'a>>);

impl<'a> Deref for BookComments<'a> {
    type Target = BTreeMap<Uuid, BookComment<'a>>;

    fn deref(&self) -> &Self::Target {
        &self.0
    }
}

impl<'a> FromIterator<&'a CommentInfo> for BookComments<'a> {
    fn from_iter<T: IntoIterator<Item=&'a CommentInfo>>(iter: T) -> Self {
        let comments = iter
            .into_iter()
            .chunk_by(|c| c.response_to)
            .into_iter()
            .sorted_unstable_by(|(a_response_to, _), (b_response_to, _)| {
                match (a_response_to, b_response_to) {
                    (None, None) => Ordering::Equal,
                    (None, _) => Ordering::Less,
                    (_, None) => Ordering::Greater,
                    (Some(a_response_to), Some(b_response_to)) => a_response_to.cmp(b_response_to),
                }
            })
            .fold(BTreeMap::new(), |mut comments, (response_to, responses)| {
                match response_to {
                    None => comments.extend(responses.map(
                        |CommentInfo {
                             id, username, text, ..
                         }| {
                            (
                                *id,
                                BookComment {
                                    username,
                                    text,
                                    responses: Self::default(),
                                },
                            )
                        },
                    )),
                    Some(response_to) => {
                        if insert_responses(&mut comments, responses, response_to).is_some() {
                            unreachable!("could not store responses")
                        }
                    }
                }

                comments
            });

        Self(comments)
    }
}

fn insert_responses<'a, I: Iterator<Item=&'a CommentInfo>>(
    comments: &mut BTreeMap<Uuid, BookComment<'a>>,
    mut responses: I,
    response_to: Uuid,
) -> Option<I> {
    match comments.entry(response_to) {
        Entry::Vacant(_) => {
            for comment in comments.values_mut() {
                match insert_responses(&mut comment.responses.0, responses, response_to) {
                    Some(r) => {
                        responses = r;
                        continue;
                    }
                    None => return None,
                }
            }
            Some(responses)
        }
        Entry::Occupied(mut response_to) => {
            response_to.get_mut().responses.0.extend(responses.map(
                |CommentInfo {
                     id, username, text, ..
                 }| {
                    (
                        *id,
                        BookComment {
                            username,
                            text,
                            responses: BookComments::default(),
                        },
                    )
                },
            ));
            None
        }
    }
}

#[cfg(test)]
mod tests {
    use pretty_assertions::{assert_eq, assert_str_eq};
    use std::iter;
    use std::sync::LazyLock;

    use super::*;

    static COMMENTS: LazyLock<(Vec<Uuid>, Vec<CommentInfo>)> = LazyLock::new(|| {
        let ids = Vec::from_iter(iter::from_fn(|| Some(Uuid::now_v7())).take(6));
        let comments = {
            vec![
                CommentInfo {
                    id: ids[0],
                    response_to: None,
                    username: "0".to_string(),
                    text: "".to_string(),
                },
                CommentInfo {
                    id: ids[1],
                    response_to: None,
                    username: "1".to_string(),
                    text: "".to_string(),
                },
                CommentInfo {
                    id: ids[2],
                    response_to: Some(ids[1]),
                    username: "2".to_string(),
                    text: "".to_string(),
                },
                CommentInfo {
                    id: ids[3],
                    response_to: Some(ids[2]),
                    username: "3".to_string(),
                    text: "".to_string(),
                },
                CommentInfo {
                    id: ids[4],
                    response_to: None,
                    username: "4".to_string(),
                    text: "".to_string(),
                },
                CommentInfo {
                    id: ids[5],
                    response_to: Some(ids[1]),
                    username: "5".to_string(),
                    text: "".to_string(),
                },
            ]
        };

        (ids, comments)
    });

    #[test]
    fn collect_book_comments() {
        let ids = &*COMMENTS.0;
        let result: BookComments<'_> = COMMENTS.1.iter().collect();

        assert_eq!(
            result,
            BookComments({
                BTreeMap::from([
                    (
                        ids[0],
                        BookComment {
                            username: "0",
                            text: "",
                            responses: Default::default(),
                        },
                    ),
                    (
                        ids[1],
                        BookComment {
                            username: "1",
                            text: "",
                            responses: BookComments(BTreeMap::from([
                                (
                                    ids[2],
                                    BookComment {
                                        username: "2",
                                        text: "",
                                        responses: BookComments(BTreeMap::from([(
                                            ids[3],
                                            BookComment {
                                                username: "3",
                                                text: "",
                                                responses: Default::default(),
                                            },
                                        )])),
                                    },
                                ),
                                (
                                    ids[5],
                                    BookComment {
                                        username: "5",
                                        text: "",
                                        responses: Default::default(),
                                    },
                                ),
                            ])),
                        },
                    ),
                    (
                        ids[4],
                        BookComment {
                            username: "4",
                            text: "",
                            responses: Default::default(),
                        },
                    ),
                ])
            })
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
