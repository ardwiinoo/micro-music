CREATE TABLE IF NOT EXISTS playlist_songs (
    playlist_id VARCHAR(50) NOT NULL,
    song_id VARCHAR(50) NOT NULL,
    PRIMARY KEY (playlist_id, song_id),
    FOREIGN KEY (playlist_id) REFERENCES playlists(id),
    FOREIGN KEY (song_id) REFERENCES songs(id)
);