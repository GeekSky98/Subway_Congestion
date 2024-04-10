use diesel::prelude::*;
use crate::db::DbConn;
use crate::models::Station;
use crate::schema::stations::dsl::*;

pub async fn calculate_congestion(conn: &DbConn, line_id: i32) -> Result<f64, diesel::result::Error> {
    let station_counts = conn.run(move |c| {
        stations.filter(line_id.eq(line_id))
            .select(train_count)
            .load::<i32>(c)
    }).await?;

    if station_counts.is_empty() {
        return Ok(0.0);
    }

    let total: i32 = station_counts.iter().sum();
    let avg_congestion = total as f64 / station_counts.len() as f64;

    Ok(avg_congestion)
}
