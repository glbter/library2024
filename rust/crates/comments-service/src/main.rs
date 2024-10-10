use axum::routing::get;
use axum::Router;
use color_eyre::Result;
use comments_service::handler;
use tracing_subscriber::util::SubscriberInitExt;

#[tokio::main]
async fn main() -> Result<()> {
    color_eyre::install()?;
    tracing_subscriber::fmt().pretty().try_init()?;

    let app = Router::new().route(
        "/books/:book_id/comments",
        get(handler::fetch_comments_for_book),
    );

    let listener = tokio::net::TcpListener::bind("0.0.0.0:3000").await?;
    axum::serve(listener, app).await?;

    Ok(())
}
