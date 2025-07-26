# Relay (Go based URL Shortener)

Relay is a high performance URL shortener built with Go and PostgreSQL.

## Architecture

The application accepts a long URL and performs a single `INSERT` into the database.

The unique short code is created directly within PostgreSQL using a `GENERATED ALWAYS AS` column. This approach is highly efficient.

The generation process is:
  1.  A new row gets a sequential `id`.
  2.  The `short_code` column calls a custom SQL function, `encode_base62(id)`.
  3.  The function shuffles the `id` to make it non-sequential and then encodes it into a base-62 string.

## Local Performance Benchmarks (Apple Silicon - M1)
```bash
adityaraj@macair testing % k6 run test.js

         /\      Grafana   /‾‾/  
    /\  /  \     |\  __   /  /   
   /  \/    \    | |/ /  /   ‾‾\ 
  /          \   |   (  |  (‾)  |
 / __________ \  |_|\_\  \_____/ 

     execution: local
        script: test.js
        output: -

     scenarios: (100.00%) 1 scenario, 3000 max VUs, 1m20s max duration (incl. graceful stop):
              * default: Up to 3000 looping VUs for 50s over 3 stages (gracefulRampDown: 30s, gracefulStop: 30s)


  █ TOTAL RESULTS 

    checks_total.......................: 612087  12241.764484/s
    checks_succeeded...................: 100.00% 612087 out of 612087
    checks_failed......................: 0.00%   0 out of 612087

    ✓ status is 201

    HTTP
    http_req_duration.......................................................: avg=196.29ms min=150µs    med=228.47ms max=333.98ms p(90)=275.51ms p(95)=294.62ms
      { expected_response:true }............................................: avg=196.29ms min=150µs    med=228.47ms max=333.98ms p(90)=275.51ms p(95)=294.62ms
    http_req_failed.........................................................: 0.00%  0 out of 612087
    http_reqs...............................................................: 612087 12241.764484/s

    EXECUTION
    iteration_duration......................................................: avg=196.35ms min=179.79µs med=228.52ms max=334.04ms p(90)=275.57ms p(95)=294.68ms
    iterations..............................................................: 612087 12241.764484/s
    vus.....................................................................: 49     min=49          max=3000
    vus_max.................................................................: 3000   min=3000        max=3000

    NETWORK
    data_received...........................................................: 94 MB  1.9 MB/s
    data_sent...............................................................: 118 MB 2.4 MB/s



running (0m50.0s), 0000/3000 VUs, 612087 complete and 0 interrupted iterations
default ✓ [======================================] 0000/3000 VUs  50s
```