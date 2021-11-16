<?php
$symbol = $_GET['s'] ?? $s = "SH605001";
$labels = $data = $symbols = [];

try {
    $dbh = new PDO('mysql:host=127.0.0.1:3306;dbname=stock', "root", "root", [PDO::MYSQL_ATTR_INIT_COMMAND=>"set names utf8"]);

    $sql = "SELECT * FROM symbol WHERE 1 ORDER BY id";

    foreach($dbh->query($sql) as $row) {
		$symbols[] = [
			"symbol" => $row['symbol'],
			"name" => $row['name'],
		];
    }

    $sql = "SELECT id, comment_count, current, exec_at FROM quote WHERE symbol = '" . $symbol .  "'  ORDER BY exec_at";

    foreach($dbh->query($sql) as $row) {
    	$labels[] = $row['exec_at'];
		$data[] = [
			"x" => $row['exec_at'],
			"comment" => $row['comment_count'],
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
    <title>wateree-添加地址</title>
	<script src="https://cdn.jsdelivr.net/npm/jquery@1.12.4/dist/jquery.min.js"></script>
	<script src="https://cdn.jsdelivr.net/npm/chart.js"></script>
	<script>
		const labels = <?php echo json_encode($labels);?>;
		const data = <?php echo json_encode($data);?>;
		const config = {
		    type: 'line',
		    data: {
		        labels: labels,
		        datasets: [{
		            label: '评论数',
		            data: data,
		            backgroundColor: [
	                'rgba(54, 162, 235, 0.2)',
	            	],
		            parsing: {
		                yAxisKey: 'comment'
		            }
		        }, {
		            label: '价格',
		            data: data,
		            backgroundColor: [
	                'rgba(255, 99, 132, 1)',
	            	],
		            parsing: {
		                yAxisKey: 'current'
		            }
		        }]
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
			echo '<li><a href="a.php?s='.$value['symbol'] . '">' . $value['name'] . '</a></li>';
		}
		?>
		</ul>
	</div>
<div>
  <canvas id="myChart"></canvas>
</div>
</body>
</html>
