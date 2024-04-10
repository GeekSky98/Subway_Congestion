#[macro_use] extern crate rocket;
use crate::db::DbConn;
use rocket::serde::{json::Json, Serialize};
use crate::functions::calculate_congestion;
use crate::models::Station;

#[derive(Serialize)]
struct StationInfo {
    station: Station,
    avg_congestion: f64,
}

#[get("/congestion/<line_id>")]
pub async fn congestion_route(line_id: i32, conn: DbConn) -> Result<Json<Vec<StationInfo>>, String> {
    let stations_info = conn.run(move |c| {
        use crate::schema::stations::dsl::*;
        use diesel::prelude::*;

        stations.filter(line_id.eq(line_id))
            .load::<Station>(c)
            .map_err(|err| err.to_string())
    }).await?;

    let avg_congestion = calculate_congestion(&conn, line_id).await.map_err(|err| err.to_string())?;

    let response: Vec<StationInfo> = stations_info.into_iter().map(|station| StationInfo {
        station,
        avg_congestion,
    }).collect();

    Ok(Json(response))
}
