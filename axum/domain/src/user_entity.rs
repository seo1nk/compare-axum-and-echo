use chrono::{DateTime, Local};

/// User Entity
#[derive(Debug, Clone, Eq, PartialEq)]
pub struct UserEntity {
    id: i32,
    name: String,
    create_at: DateTime<Local>,
}
