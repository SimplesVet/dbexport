package dbexport

// SET global general_log = 1;
// SET global log_output = 'table';
// select * from mysql.general_log where event_time >= CURRENT_TIMESTAMP() - interval 10 SECOND;

// check if database has the correct configuration
// select general_log
// check actions in returned rows
// search for key words DROP, CREATE
// call specif actions from dbexport
