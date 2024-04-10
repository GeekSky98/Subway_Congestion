use rocket_sync_db_pools::{database, diesel};

#[database("postgres_db")]
pub struct DbConn(diesel::PgConnection);

pub fn stage() -> AdHoc {
    AdHoc::on_ignite("Database Migrations", |rocket| async {
        rocket.attach(DbConn::fairing())
    })
}
