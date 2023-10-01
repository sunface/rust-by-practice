use serde::Serialize;
use std::convert::Infallible;
use thiserror::Error;
use warp::{http::StatusCode, Rejection, Reply};

#[derive(Error, Debug)]
pub enum Error {
    #[error("Wrong credentials")]
    WrongCredentialsError,
    #[error("JWT token not valid")]
    JWTTokenError,
    #[error("JWT token creation error")]
    JWTTokenCreationError,
    #[error("No auth header")]
    NoAuthHeaderError,
    #[error("Invalid auth header")]
    InvalidAuthHeaderError,
    #[error("No permission")]
    NoPermissionError,
}

#[derive(Serialize, Debug)]
struct ErrorResponse {
    message: String,
    status: String,
}

impl warp::reject::Reject for Error {}

pub async fn handler_rejection(err: Rejection) -> std::result::Result<impl Reply, Infallible> {
    let (code, message) = if err.is_not_found() {
        (StatusCode::NOT_FOUND, "Not found".to_string())
    } else if let Some(e) = err.find::<Error>() {
        match e {
            Error::WrongCredentialsError => (StatusCode::FORBIDDEN, e.to_string()),
            Error::NoPermissionError => (StatusCode::UNAUTHORIZED, e.to_string()),
            Error::JWTTokenError => (StatusCode::UNAUTHORIZED, e.to_string()),
            Error::JWTTokenCreationError => (StatusCode::INTERNAL_SERVER_ERROR, e.to_string()),
            _ => (StatusCode::BAD_REQUEST, e.to_string()),
        }
    } else if err.find::<warp::reject::MethodNotAllowed>().is_some() {
        (
            StatusCode::METHOD_NOT_ALLOWED,
            "Method Not Allowed".to_string(),
        )
    } else {
        eprint!("Unhandled Error {:?}", err);
        (
            StatusCode::INTERNAL_SERVER_ERROR,
            "Internal server error".to_string(),
        )
    };
    let json = warp::reply::json(&ErrorResponse {
        status: code.to_string(),
        message,
    });
    Ok(warp::reply::with_status(json, code))
}
