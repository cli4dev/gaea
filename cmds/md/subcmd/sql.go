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

// GetAllTableNameInOracle .
const GetAllTableNameInOracle = `
select
  ub.table_name
from
  user_tables ub
order by 
  ub.table_name
`

// GetSingleTableInfoInOracle .
const GetSingleTableInfoInOracle = `
select
    a.column_name,
    a.data_type,
    case upper(a.data_type)
    when upper('number') then
    decode(a.data_scale,
            0,
            to_char(a.data_precision),
            a.data_precision || ',' || a.data_scale)
    when upper('date') then
      ''
    else
     to_char(a.data_length)
    end data_length,
    a.data_precision,
    a.data_scale,
    a.nullable,
    long_to_char(a.table_name,a.column_id) data_default,
    b.comments column_comments,
    d.comments table_comments
from
    user_tab_columns a,
    user_col_comments b,
    (select column_name from user_ind_columns where table_name = @table_name) c,
    user_tab_comments d
where
    a.table_name = b.table_name
    and a.table_name = d.table_name
    and a.column_name = b.column_name
    and a.column_name = c.column_name(+)
    and a.table_name = @table_name
order by 
    a.column_id
`
