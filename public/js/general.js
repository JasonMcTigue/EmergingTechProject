	$(document).ready(function(){
	//initial
	$('#templates').load('templates/index.html');
	alert('hello');
	//handle menu clicks
	$('div div nav div div#navbar ul li a').click(function() {
		alert('hello again');
		
		var page = $(this).attr('href');
		$('#templates').load('templates/' + page + '.html');
		return false;
	});
});