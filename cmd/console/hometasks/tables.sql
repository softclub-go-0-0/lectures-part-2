create table teachers
(
    id         bigint generated always as identity primary key,
    name       text not null,
    surname    text not null,
    phone      text not null,
    email      text,
    created_at timestamptz default now(),
    updated_at timestamptz,
    deleted_at timestamptz
);

create table timetables
(
    id         bigint generated always as identity primary key,
    classroom  text not null,
    start      time not null,
    finish     time not null,
    created_at timestamptz default now(),
    updated_at timestamptz,
    deleted_at timestamptz
);

create table courses
(
    id          bigint generated always as identity primary key,
    title       text   not null,
    monthly_fee bigint not null,
    duration    bigint not null,
    created_at  timestamptz default now(),
    updated_at  timestamptz,
    deleted_at  timestamptz
);

create table groups
(
    id           bigint generated always as identity primary key,
    course_id    bigint references courses (id) not null,
    teacher_id   bigint references teachers (id),
    timetable_id bigint references timetables (id),
    title        text                           not null,
    start_date   date,
    created_at   timestamptz default now(),
    updated_at   timestamptz,
    deleted_at   timestamptz
    --constraint "luboy" foreign key (teacher_id) references teachers(id)
);

create table students
(
    id         bigint generated always as identity primary key,
    group_id   bigint references groups (id),
    name       text not null,
    surname    text not null,
    phone      text not null,
    email      text,
    created_at timestamptz default now(),
    updated_at timestamptz,
    deleted_at timestamptz
);