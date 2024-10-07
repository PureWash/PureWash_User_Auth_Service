CREATE Table company (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100) NOT NULL,
    description TEXT,
    created_at TIMESTAMP default CURRENT_TIMESTAMP,
    updated_at TIMESTAMP default CURRENT_TIMESTAMP,
    deleted_at int DEFAULT 1     --0 bo'lsa o'chirilgan bo'ladi
);

CREATE Table services (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    tariffs  text,
    name VARCHAR(100) not null,
    description TEXT,
    price DECIMAL(10,2)
);


CREATE Table orders (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    client_id UUID NOT NULL REFERENCES clients(id) ON DELETE CASCADE,
    service_id UUID NOT NULL REFERENCES services(id) ON DELETE CASCADE,
    area float NOT NULL,
    total_price float NOT NULL,
    status VARCHAR(50) DEFAULT 'inprogress',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at int DEFAULT 1
);