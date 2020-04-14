<?php

// 컨텐츠 타입을 json으로 지정합니다.

header("Content-Type: application/json");



// 검색될 샘플 데이터 입니다.

$cities = array("서울","부산","대구","광주","울산");



// 넘어온 검색어 파라미터 입니다.

$term = $_GET['term'];



// 데이터를 루핑 하면서 찾습니다.

$result = [];

foreach($cities as $city) {

    if(strpos($city, $term) !== false) {

        $result[] = array("label" => $city, "value" => $city);

    }

}



// 찾아진 데이터를 json 데이터로 변환하여 전송합니다.

echo json_encode($result);

?>