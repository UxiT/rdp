CREATE TABLE IF NOT EXISTS 'user_courses' (
    user_id int,
    cource_id int,
    PRIMARY KEY (user_id, cource_id),
    FOREIGN KEY (user_id) references users(id) on delete cascade,
    FOREIGN KEY (cource_id) references cources(id) on delete cascade,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
)

CREATE TABLE IF NOT EXISTS 'groups' (
    id SERIAL PRIMARY KEY,
    name char(256),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
)

CREATE TABLE IF NOT EXISTS 'groups_courses' (
    group_id int,
    cource_id int,
    PRIMARY KEY (group_id, cource_id),
    FOREIGN KEY (group_id) references groups(id) on delete cascade,
    FOREIGN KEY (cource_id) references cources(id) on delete cascade,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
)

CREATE TABLE IF NOT EXISTS 'students' (
    id SERIAL PRIMARY KEY,
    user_id int,
    group_id int,
    FOREIGN KEY (user_id) references users(id) on delete cascade,
    FOREIGN KEY (group_id) references groups(id),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
)

CREATE TABLE IF NOT EXISTS 'teachers' (
    id SERIAL PRIMARY KEY,
    user_id int,
    FOREIGN KEY (user_id) references users(id) on delete cascade,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
)

CREATE TABLE IF NOT EXISTS 'teachers_groups' (
    teacher_id int,
    group_id int,
    PRIMARY KEY (group_id, teacher_id),
    FOREIGN KEY (teacher_id) references teachers(id) on delete cascade,
    FOREIGN KEY (group_id) references groups(id) on delete cascade,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
)

CREATE TABLE IF NOT EXISTS 'tasks' (
    id SERIAL PRIMARY KEY,
    name char(256) NOT NULL,
    cource_id int NOT NULL,
    theory_file char(256),
    rdp_config char(256),
    description text NOT NULL,
    extra_file JSON
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
)