CREATE Table company (
                         id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                         name VARCHAR(100) NOT NULL,
                         description TEXT,
                         logo_url VARCHAR,
                         created_at TIMESTAMP default CURRENT_TIMESTAMP,
                         updated_at TIMESTAMP,
                         deleted_at int DEFAULT 1     --0 bo'lsa o'chirilgan bo'ladi
);

CREATE Table services (
                          id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                          tariffs  text,
                          name VARCHAR(100) not null,
                          description TEXT,
                          price DECIMAL(10,2)
);

CREATE Table addresses (
                           id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                           user_id UUID ,
                           latitude VARCHAR(100),
                           longitude VARCHAR(100),
                           created_at TIMESTAMP default CURRENT_TIMESTAMP,
                           updated_at TIMESTAMP,
                           deleted_at int DEFAULT 1
);

CREATE Table orders (
                        id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                        user_id UUID ,
                        service_id UUID REFERENCES services(id) ON DELETE CASCADE,
                        area float NOT NULL,
                        total_price float,
                        status VARCHAR(50),
                        created_at TIMESTAMP default CURRENT_TIMESTAMP,
                        updated_at TIMESTAMP,
                        deleted_at int DEFAULT 1
);