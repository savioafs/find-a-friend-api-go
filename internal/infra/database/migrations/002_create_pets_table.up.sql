CREATE TABLE pets (
    id VARCHAR(255) PRIMARY KEY,
    org_id VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    about TEXT,
    age VARCHAR(50),
    size VARCHAR(50),
    energy_level VARCHAR(50),
    dependency_level VARCHAR(50),
    enviroment VARCHAR(255),
    photos JSON,
    requirements_for_adoption JSON,
    FOREIGN KEY (org_id) REFERENCES organizations (id) ON DELETE CASCADE
);
