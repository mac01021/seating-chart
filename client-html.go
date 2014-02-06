package main


var ClientHTML = `
<!DOCTYPE html>
<html ng-app>
<head>
	<script type="text/javascript"
		src="http://ajax.googleapis.com/ajax/libs/angularjs/1.1.5/angular.min.js">
	</script>
	<script type="text/javascript"
		src="client.js">
	</script>
	<style type="text/css">
		.centered {
			margin-left:auto;
			margin-right:auto;
			width:50%;
			background-color:#8b4513;
			border-radius: 15px;
		}

		input.filt{
			-moz-border-radius: 15px;
			border-radius: 15px;
			border: solid 1px black;
			padding: 5px;
		}

		.odd {
			background-color:#ae7846;
		}

	</style>
</head>




<body ng-controller="EventHandler">
    <br><br>
    <br><br>
    <div class="centered" >
	<br>
	<table class="centered">
		<thead>
			<tr>
				<td>
					<input class="filt" type="text" ng-model="Filter.name" placeholder="Name" />
				</td>
				<td colspan="2">
					<input class="filt" type="text" ng-model="Filter.table" placeholder="Table" />
				</td>
			</tr>
		</thead>
		<tbody>
			<tr ng-repeat="guest in Chart | filter: Filter | orderBy: guest.name" >
				<td ng-class-even="'even'" ng-class-odd="'odd'" >
					{{guest.name}}
				</td>
				<td ng-class-even="'even'" ng-class-odd="'odd'" >
					&nbsp;&nbsp;&nbsp;&nbsp;{{guest.table}}
				</td>
				<td ng-class-even="'even'" ng-class-odd="'odd'" >
					<input name="{{guest.name}}"
						type="radio"
						ng-model="Arrived[guest.name]"
						ng-value="true"
						ng-change="Report(guest.name)"/>
				</td>
			</tr>
		</tbody>
	</table>
	<br><br>
	&nbsp;&nbsp;&nbsp;&nbsp;{{ Status }}
	<br>
    </div>
</body>
</html>
`



