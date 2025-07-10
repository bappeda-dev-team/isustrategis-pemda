CREATE TABLE tb_data_terukur_pendukung_pernyataan (
    id INT PRIMARY KEY AUTO_INCREMENT,
    alasan_kondisi_id INT NOT NULL,
    data_terukur TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    CONSTRAINT fk_tb_alasan_kondisi FOREIGN KEY (alasan_kondisi_id)
        REFERENCES tb_alasan_kondisi(id)
        ON DELETE CASCADE
        ON UPDATE CASCADE
) ENGINE=InnoDB;