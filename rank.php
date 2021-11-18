<?php
$key = $_GET['key'] ?? "comment";
$labels = $data = $symbols = [];

try {
    $dbh = new PDO('mysql:host=127.0.0.1:3306;dbname=stock', "root", "root", [PDO::MYSQL_ATTR_INIT_COMMAND=>"set names utf8"]);

    $sql = "SELECT * FROM symbol WHERE 1 ORDER BY id";

    foreach($dbh->query($sql) as $row) {
        $symbols[$row['symbol']] = [
            "id" => $row['id'],
            "symbol" => $row['symbol'],
            "name" => $row['name'],
            "background_color" => $row['background_color'],
            "border_color" => $row['border_color'],
        ];
    }

    $sql = "SELECT id,symbol, comment_count,comment_count3, current,low,high, exec_at FROM quote WHERE 1  ORDER BY exec_at";

    foreach($dbh->query($sql) as $row) {
        $labels[$row['exec_at']] = $row['exec_at'];
        $data[$symbols[$row['symbol']]['id']]['name'] = $symbols[$row['symbol']]['name'] ?? $row['symbol'];
        $data[$symbols[$row['symbol']]['id']]['backgroundColor'] = $symbols[$row['symbol']]['background_color'] ?? "rgba(255, 99, 132, 1)";
        $data[$symbols[$row['symbol']]['id']]['borderColor'] = $symbols[$row['symbol']]['border_color'] ?? "rgba(255, 99, 132,0.5)";
        $data[$symbols[$row['symbol']]['id']]['data'][] = [
            "x" => $row['exec_at'],
            "low" => $row['low'],
            "high" => $row['high'],
            "comment" => $row['comment_count'],
            "comment3" => $row['comment_count3'],
            "current" => $row['current'],
        ];
    }
    ksort($data);
    $datasets = [];
    foreach ($data as $value) {
    	$datasets[] = [
    		'label' => $value["name"],
    		'data' => $value['data'] ,
    		'backgroundColor' => $value["backgroundColor"],
    		'borderColor' =>  $value["borderColor"],
    		'parsing' => [
    			'yAxisKey' => $key,
    			'yAxisID' => $key,
    		]];
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
		const labels = <?php echo json_encode(array_values($labels));?>;
		const datasets = <?php echo json_encode($datasets)?>;

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
                    text: '热门股票评论排行'
                  }
                }
            },
		    data: {
		        labels: labels,
		        tension: 0.1,
		        datasets: datasets,
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
		foreach ($symbols as $value) {
			echo '<li><a href="index.php?s='.$value['symbol'] . '">' . $value['name'] . '</a></li>';
		}
		?>
		<?php if($key == "comment"):?>
		    <li><a href="/rank.php?key=comment3">机构评论</a></li>
		<?php else:?>
			<li><a href="/rank.php?key=comment">热门评论</a></li>
		<?php endif;?>
		</ul>

	</div>
<div>
  <canvas id="myChart"></canvas>
</div>
</body>
</html>