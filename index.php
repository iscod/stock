<?php
$symbol = $_GET['s'] ?? $s = "SZ000651";
$labels = $data = $symbols = [];

try {
    $dbh = new PDO('mysql:host=127.0.0.1:3306;dbname=stock', "root", "root", [PDO::MYSQL_ATTR_INIT_COMMAND=>"set names utf8"]);

    $sql = "SELECT * FROM symbol WHERE 1 ORDER BY id";

    foreach($dbh->query($sql) as $row) {
		$symbols[$row['symbol']] = [
			"symbol" => $row['symbol'],
			"name" => $row['name'],
		];
    }

    $sql = "SELECT id, comment_count,comment_count3, current,low,high,open, volume, summary_volume,fund_flow, exec_at FROM quote WHERE symbol = '" . $symbol .  "'  ORDER BY exec_at";

    foreach($dbh->query($sql) as $row) {
        $labels[] = $row['exec_at'];
        $data[] = [
            "x" => $row['exec_at'],
            "open" => $row['open'],
            "low" => $row['low'],
            "high" => $row['high'],
            "volume" => $row['volume']/100,
            "summary_volume" => json_decode($row['summary_volume'], true),
            "fund_flow" => json_decode($row['fund_flow'], true),
            "comment" => $row['comment_count'],
            "comment3" => $row['comment_count3'],
            "current" => $row['current'],
        ];
    }
    $dbh = null;
} catch (PDOException $e) {
    print "Error!: " . $e->getMessage() . "<br/>";
    die();
}

?>
<!doctype html>
<html lang="zh-CN">
<head>
    <meta charset="utf-8">
    <meta name="viewport" content="initial-scale=1.0, user-scalable=no, width=device-width">
    <title>AiStock</title>
	<script src="https://cdn.jsdelivr.net/npm/jquery@1.12.4/dist/jquery.min.js"></script>
	<script src="https://cdn.jsdelivr.net/npm/chart.js"></script>
	<script>
		const labels = <?php echo json_encode($labels);?>;
		const data = <?php echo json_encode($data);?>;
		const config = {
		    type: 'line',
		    options: {
                responsive: true,
                plugins: {
                  legend: {
                    position: 'top',
                  },
                  title: {
                    display: true,
                    text: '<?php echo $symbols[$symbol]["name"] ?? $symbol ?>'
                  }
                }
            },
            data: {
                labels: labels,
                tension: 0.1,
                datasets: [
                    {
                        label: '价格',
                        data: data,
                        backgroundColor: 'rgba(255, 99, 132, 1)',
                        borderColor: 'rgba(255, 99, 132, 1)',
                        borderWidth: 1,
                        parsing: {
                            yAxisKey: 'current',
                        },
                        yAxisID: 'current',
                    },
                    {
                        label: '最低价',
                        data: data,
                        backgroundColor: 'rgba(54, 162, 235, 0.1)',
                        borderColor: 'rgba(54, 162, 235, 0.2)',
                        borderWidth: 1,
                        parsing: {
                            yAxisKey: 'low'
                        },
                        yAxisID: 'current',
                        fill: '+1'
                    },
                    {
                        label: '最高价',
                        data: data,
                        backgroundColor: 'rgba(0, 128, 0, 0.1)',
                        borderColor: 'rgba(0, 128, 0, 0.2)',
                        borderWidth: 1,
                        parsing: {
                            yAxisKey: 'high'
                        },
                        yAxisID: 'current',
                    },

                    {
                        label: '评论数',
                        data: data,
                        backgroundColor: 'rgba(54, 162, 235, 1)',
                        borderColor: 'rgba(54, 162, 235, 1)',
                        borderWidth: 1,
                        parsing: {
                            yAxisKey: 'comment',
                        },
                        yAxisID: 'comment',
                    },
                    {
                        label: '机构',
                        data: data,
                        backgroundColor: 'rgba(54, 162, 235, 0.5)',
                        borderColor: 'rgba(54, 162, 235, 0.5)',
                        borderWidth: 1,
                        parsing: {
                            yAxisKey: 'comment3',
                        },
                        yAxisID: 'comment',
                    },
                    {
                        label: '成交量(手)',
                        data: data,
                        backgroundColor: 'rgba(76, 175, 80, 0.5)',
                        borderColor: 'rgba(76, 175, 80, 0.5)',
                        borderWidth: 1,
                        parsing: {
                            yAxisKey: 'volume',
                        },
                        yAxisID: 'volume',
                    },
                    {
                        label: '卖盘',
                        data: data,
                        backgroundColor: 'rgba(153, 102, 255, 0.2)',
                        borderColor: 'rgba(153, 102, 255, 0.2)',
                        borderWidth: 1,
                        parsing: {
                            yAxisKey: 'summary_volume.10.s',
                        },
                        yAxisID: 'volume',
                    },
                    {
                        label: '买盘',
                        data: data,
                        backgroundColor: 'rgba(75, 192, 192, 0.2)',
                        borderColor: 'rgba(75, 192, 192, 0.2)',
                        borderWidth: 1,
                        parsing: {
                            yAxisKey: 'summary_volume.10.b',
                        },
                        yAxisID: 'volume',
                    },
                    {
                      label: '中盘',
                      data: data,
                      backgroundColor: 'rgba(205, 220, 57, 0.4)',
                      borderColor: 'rgba(205, 220, 57, 0.4)',
                      borderWidth: 1,
                      parsing: {
                          yAxisKey: 'summary_volume.10.m',
                      },
                      yAxisID: 'volume',
                    },
                    {
                      label: '主力净流入',
                      data: data,
                      type:"bar",
                      backgroundColor: 'rgba(255, 159, 64, 0.2)',
                      borderColor: 'rgba(255, 159, 64, 1)',
                      borderWidth: 1,
                      parsing: {

                          yAxisKey: 'fund_flow.MainNetIn',
                      },
                      yAxisID: 'fund_flow.MainNetIn',

                    },
                ],
            },
        };

		$(document).ready(function () {
		    const myChart = new Chart(
		    	document.getElementById('myChart'),
		    	config
		  	);
		});
	</script>
    <style>
    	a {
    		font-size: 12px;
    		color: #555;
    		text-decoration: none;
    	}
    	a:hover{
    		    color: #337ab7;
    	}
    	ul li {
    		padding: 2px;
    		list-style: none;
    		display: inline-block;
    	}
	</style>
</head>
<body>
	<div>
		<ul>
		<?php
		foreach ($symbols as $key => $value) {
			echo '<li><a href="index.php?s='.$value['symbol'] . '">' . $value['name'] . '</a></li>';
		}
		?>
		<li><a href="/rank.php">评论排行</a></li>
		</ul>
	</div>
<div>
  <canvas id="myChart"></canvas>
</div>
</body>
</html>
