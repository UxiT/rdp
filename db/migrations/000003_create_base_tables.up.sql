CREATE TABLE IF NOT EXISTS user_courses (
    user_id int,
    course_id int,
    PRIMARY KEY (user_id, course_id),
    FOREIGN KEY (user_id) references users(id) on delete cascade,
    FOREIGN KEY (course_id) references courses(id) on delete cascade,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS groups (
    id SERIAL PRIMARY KEY,
    title char(256),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS groups_courses (
    group_id int,
    course_id int,
    PRIMARY KEY (group_id, course_id),
    FOREIGN KEY (group_id) references groups(id) on delete cascade,
    FOREIGN KEY (course_id) references courses(id) on delete cascade,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS students (
    id SERIAL PRIMARY KEY,
    user_id int,
    group_id int,
    FOREIGN KEY (user_id) references users(id) on delete cascade,
    FOREIGN KEY (group_id) references groups(id),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS teachers (
    id SERIAL PRIMARY KEY,
    user_id int,
    FOREIGN KEY (user_id) references users(id) on delete cascade,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS teachers_groups (
    teacher_id int,
    group_id int,
    PRIMARY KEY (group_id, teacher_id),
    FOREIGN KEY (teacher_id) references teachers(id) on delete cascade,
    FOREIGN KEY (group_id) references groups(id) on delete cascade,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS tasks (
    id SERIAL PRIMARY KEY,
    title char(256) NOT NULL,
    course_id int NOT NULL,
    theory_file char(256),
    rdp_config char(256),
    description text NOT NULL,
    extra_file JSON,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS users_tasks (
    student_id int not null,
    task_id int not null,
    PRIMARY KEY (student_id, task_id),
    FOREIGN KEY (student_id) references students(id) on delete cascade,
    FOREIGN KEY (task_id) references tasks(id) on delete cascade,
    score int not null DEFAULT 0,
    max_score int not null DEFAULT 10,
    status int not null DEFAULT 0,
    report char(256),
    comment char(256),

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);