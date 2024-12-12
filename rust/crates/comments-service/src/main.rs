use std::{env, net::Ipv4Addr};

use clap::Parser;
use color_eyre::Result;
use comments_service::{
    repo::{CommentRepoImpl, SessionRepoImpl},
    ProdAppState,
};
use sqlx::PgPool;
use tokio::net::TcpListener;
use tracing_subscriber::util::SubscriberInitExt;

#[tokio::main]
async fn main() -> Result<()> {
    let Args { port, lazy } = Args::parse();

    color_eyre::install()?;
    tracing_subscriber::fmt().pretty().finish().try_init()?;

    let session_cookie_name = env::var("SESSION_COOKIE").unwrap_or("session".to_owned());
    let db_url = env::var("DATABASE_URL")?;
    let pg_pool = if lazy {
        PgPool::connect_lazy(&db_url)?
    } else {
        PgPool::connect(&db_url).await?
    };
    let comment_repo = CommentRepoImpl::new(pg_pool.clone());
    let session_repo = SessionRepoImpl::new(pg_pool);
    let app = comments_service::build(ProdAppState::new(
        comment_repo,
        session_repo,
        session_cookie_name,
    ));

    let listener = TcpListener::bind((Ipv4Addr::new(0, 0, 0, 0), port)).await?;
    axum::serve(listener, app).await?;

    Ok(())
}

#[derive(Debug, Parser)]
#[command(about, long_about = None)]
struct Args {
    /// server port
    #[arg(short, long, default_value_t = 3000)]
    port: u16,
    /// weather to connect to the database lazily
    #[arg(long)]
    lazy: bool,
}
