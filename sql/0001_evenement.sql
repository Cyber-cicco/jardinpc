CREATE TABLE IF NOT EXISTS evenement (
    ID BIGINT PRIMARY KEY AUTO_INCREMENT,
    Title VARCHAR(60) NOT NULL,
    Description VARCHAR(512) NULL,
    Date DATETIME NOT NULL,
    DateCreation DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    Illustration VARCHAR(60) NULL,
    CreateurID BIGINT NOT NULL
);


