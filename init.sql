-- Файл для инициализации БД PostgreSQL
-- Запустите: psql -U postgres -f init.sql

-- Создаём базу данных (если её нет)
DROP DATABASE IF EXISTS volleyball;
CREATE DATABASE volleyball OWNER postgres;

-- Подключаемся к базе
\c volleyball

-- Создаём таблицу игроков
CREATE TABLE players (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    position VARCHAR(50),
    age INTEGER,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Создаём таблицу матчей
CREATE TABLE matches (
    id SERIAL PRIMARY KEY,
    home_team VARCHAR(100),
    away_team VARCHAR(100),
    home_score INTEGER,
    away_score INTEGER,
    date DATE
);

-- Создаём таблицу статистики
CREATE TABLE stats (
    id SERIAL PRIMARY KEY,
    player_id INTEGER REFERENCES players(id),
    match_id INTEGER REFERENCES matches(id),
    points INTEGER,
    assists INTEGER,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Создаём таблицу сезонов
CREATE TABLE seasons (
    id SERIAL PRIMARY KEY,
    season_name VARCHAR(50),
    start_date DATE,
    end_date DATE
);

-- Вставляем данные в таблицу сезонов
INSERT INTO seasons (season_name, start_date, end_date) VALUES
('2024/2025', '2024-09-01', '2025-06-30');
