package subcmd

const QueryMysql = `
select
   c.table_name,
   c.column_name,
   c.column_default,
   c.is_nullable,
   c.column_type,
   c.column_key,
   c.column_comment,
   t.table_comment
from
  information_schema.columns c
  inner join information_schema.tables t on c.table_name= t.table_name
where c.table_schema = @schema
`

const QueryOracle = `
select
  t .database_name as database_name,
  t .table_name as table_name,
  t .column_name as column_name,
  t .column_type as column_type,
  t .data_length as data_length,
  t .nullable as nullable,
  -- t.data_default as data_default, 
  t .column_comment as column_comment,
 -- b.constraint_type as constraint_type,
   t.table_comment as table_comment
from
  (
    select
      ub.tablespace_name as database_name,
      utc.table_name as table_name,
      utc.column_name as column_name,
      utc.data_type as column_type,
      utc.data_length as data_length,
      utc.nullable  as nullable,
      utc.data_default   as data_default,
      utcm.comments as table_comment,
      ucc.comments as column_comment
    from
      user_tables ub
    inner join user_tab_columns utc on ub.table_name = utc.table_name
    inner join user_col_comments ucc on utc.column_name = ucc.column_name
    inner join user_tab_comments utcm on ub.table_name = utcm.table_name
    and utc.table_name = ucc.table_name
  ) t
left join (
  select
    ucc.table_name as table_name,
    ucc.column_name as column_name,
    wm_concat (uc.constraint_type) as constraint_type
  from
    user_cons_columns ucc
  left join user_constraints uc on ucc.constraint_name = uc.constraint_name
  group by
    ucc.table_name,
    ucc.column_name
) b on t .table_name = b.table_name
and t .column_name = b.column_name
order by  t .table_name , t .column_name
`
