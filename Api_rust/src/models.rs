use serde::{Serialize, Deserialize};
use diesel::prelude::*;
use diesel::table;

#[derive(Serialize, Deserialize, Queryable)]
pub struct Line {
    pub line_id: i32,
    pub line_name: String,
}

#[derive(Serialize, Deserialize, Queryable)]
pub struct Station {
    pub station_id: i32,
    pub station_name: String,
    pub train_count: i32,
    pub line_id: i32,
}

table! {
    lines (line_id) {
        line_id -> Int4,
        line_name -> Varchar,
    }
}

table! {
    stations (station_id) {
        station_id -> Int4,
        station_name -> Varchar,
        train_count -> Int4,
        line_id -> Int4,
    }
}
