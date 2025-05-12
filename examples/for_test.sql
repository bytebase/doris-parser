WITH input_query AS (
  SELECT 
    CAST(arr[1] AS STRING) as id, 
    CAST(arr[2] AS STRING) as agent_code, 
    CAST(arr[3] AS STRING) as line_code, 
    CAST(arr[4] AS STRING) as game_service_id, 
    CAST(arr[5] AS DATETIME) as finish_game_date_time_start, 
    CAST(arr[6] AS DATETIME) as finish_game_date_time_end 
  FROM 
    (
      SELECT 
        split_by_string(content, ',') as arr 
      FROM 
        (
          SELECT 
            1 k1
        ) AS t LATERAL VIEW EXPLODE (
          [ '234,APICNY,024,20012001,2025-02-04 18:09:40,2025-02-05 09:32:40' ]
        ) tmp1 AS content
    ) s
), 
raw_query AS (
  SELECT 
    i.id, 
    i.agent_code, 
    i.line_code, 
    i.game_service_id, 
    d.kill_rate, 
    d.finish_game_date_time, 
    d.valid_bet, 
    d.item_count 
  FROM 
    input_query i 
    LEFT JOIN game_dwd.user_pay_income_record_details d ON i.agent_code = d.agent_code 
    AND i.line_code = d.line_code 
    AND i.game_service_id = d.game_service_id 
    AND d.finish_game_date_time >= i.finish_game_date_time_start 
    AND d.finish_game_date_time < i.finish_game_date_time_end 
    AND d.line_code != '@@@' 
    AND d.porder_type IN (1, 2)
), 
over_query AS (
  SELECT 
    id, 
    agent_code, 
    line_code, 
    game_service_id, 
    kill_rate, 
    sum(valid_bet) over (
      partition by id, agent_code, line_code, 
      game_service_id
    ) as valid_bet, 
    sum(item_count) over (
      partition by id, agent_code, line_code, 
      game_service_id
    ) as item_count, 
    rank() over (
      partition by id, 
      agent_code, 
      line_code, 
      game_service_id 
      order by 
        finish_game_date_time DESC
    ) as rank 
  FROM 
    raw_query
) 
SELECT 
  /*+ SET_VAR(enable_partition_topn = false) */
  id, 
  agent_code, 
  line_code, 
  game_service_id, 
  valid_bet, 
  item_count, 
  kill_rate AS last_kill_rate 
FROM 
  over_query 
WHERE 
  rank = 1;