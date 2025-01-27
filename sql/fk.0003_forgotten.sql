ALTER TABLE forgotten_password_attempt ADD CONSTRAINT FOREIGN KEY(UtilisateurID) REFERENCES utilisateur(ID);

