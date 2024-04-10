#[macro_use] extern crate rocket;

mod db;
mod models;
mod routes;
mod functions;

use routes::{congestion_route};

#[launch]
fn rocket() -> _ {
    rocket::build()
        .mount("/", routes![congestion_route])
        .attach(db::stage())
}
