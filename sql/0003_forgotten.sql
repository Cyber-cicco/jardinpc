CREATE TABLE forgotten_password_attempt (
    ID BIGINT PRIMARY KEY AUTO_INCREMENT,
    DateDemande DATETIME NOT NULL,
    LienChangement BINARY(16) NOT NULL,
    Active TINYINT(1) DEFAULT 1 NOT NULL,
    UtilisateurID BIGINT NOT NULL
);

