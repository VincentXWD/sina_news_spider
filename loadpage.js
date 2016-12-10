function loadpage(i) {
	m=(i-1)*count;
	n=i*count;
	if(n>(dt_i+1)) {
		n=dt_i+1;
	}
	echo(m,n);
	page(i);
	return true;
}

function echo(i,j) {
	var outstr='';
	if(j>(dt_i-1))
	j=dt_i;
	if(i<0||i>(dt_i-1))
	i=0;
	for(m=i;m<j;m++) {
		outstr+=data_p[m][0];	
	}
	document.getElementById("show_data").innerHTML=outstr;
}

function page(i) {
	var outstr='';
	if(i<=1)
	outstr+=' <span class="btn_off">&lt;&lt;Previous</span>';
	else
	outstr+=' <a href="#top" title="Previous" onclick="return loadpage('+(i-1)+');">&lt;&lt;Previous</a>';

	start_=i-2;
	end_=i+3;
	if(start_<1) {
		start_=1;
		end_=6;
	}
	if(end_>page_count) {
		end_=page_count+1;
		start_=end_-5
	}
	for(m=start_;m<end_;m++) {
		if(m>0) {
			if(m==i) outstr+='<b><span class="pagebox_cur_page">'+m+'</span></b>';
			else outstr+='<a href="#top" onclick="return loadpage('+m+');" title="">'+m+'</a>';
		}
	}
	if((i*count)>=(dt_i-1)) outstr+='<span class="btn_off">Next&gt;&gt;</span>';
	else outstr+='<a class="btn_on" href="#top" onclick="return loadpage('+(i+1)+');" title="next">Next&gt;&gt;</a>';
	document.getElementById("show_page").innerHTML=outstr;
	//document.getElementById("ast_page1").innerHTML=outstr;
}