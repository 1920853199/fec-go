Mongo������
============


### ����

1.�ֿ�ֱ�

����ʱ����зֿ⣬һ��һ����

����website_id���зֱ�һ��websiteһ����


2.�¼Ӽ����ֶΣ�

`service_timestamp`: �������������ݵ�ʱ���

`service_datetime`: ��������������, ��ʽ��Y-m-d H:i:s

`service_date`: ��������������, ��ʽ��Y-m-d

`stay_seconds`: ͨ���������ڵ�uuid��`service_timestamp`�������ã���ѯuuid = ��ǰuuid && service_date = ��ǰ��service_date

`uuid_first_page`:���ڰ���ʱ��ֿ⣬վ��ֱ���ѯ��ǰ���Ƿ����uuid����������ڣ��� uuid_first_page = 1������ uuid_first_page = 0

`ip_first_page`: �����ʱû�з����ô�

`uuid_first_category`: ������� uuid ��category�����ѯ�Ƿ���ڣ���������ڣ���Ϊ1����������Ϊ0

`ip_first_category`: �����ô�

`url_new`: ȥ��ĳЩ�������url��Ʃ�磺

```
$var_names = [
    'utmsource',
    'cid',
    'aid',
    'mid',
    'affiliate',
    'click',
    'utm_source',
    'utm_medium',
    'utm_campaign',
    'utm_content',
    '_ga',
    'gclid',
];	
	$get['url_new'] = trimVar(removeqsvar($url, $var_names));
```

`search_login_email`: 

```
            where([
				'uuid'=>$one['uuid'],
				'login_email'=>['$exists' => true]
			])
```


1.
```
function removeqsvar($url, $var_names) {
	foreach($var_names as $param){
		$url = preg_replace('/([?&])'.$param.'=[^&]+(&|$)/','$1',$url);
	}
	return $url;
}
```

2.ȥ�������ַ���ȥ����Լ350���ַ�

```
function trimVar($url){
	$url = trim($url,"?");
	$url = trim($url,"#");
	$url = trim($url,"&");
	$url = trim($url,"/");
	$url_length = strlen($url);
	# ���url��Լ350����ôֻȡ350�ڵ�url�ַ�
	if($url_length > 350){
		$url = substr($url,0,350);
	}
	return $url;
}
```


`first_visit_this_url`��������� uuid ��category

������

`order.invoice`: �����ڽ������ݲ��֣����¶�����֧��״̬

`uuid, service_date, _id`: ���ڼ��� stay_seconds ��ֵ����ѯuuid = ��ǰuuid && service_date = ��ǰ��service_date�� _id���� ��ȡһ��ֵ��Ȼ��鿴ʱ����





��Ҫ�ӵ�����

1.����website�ֱ�

������
order.invoice: