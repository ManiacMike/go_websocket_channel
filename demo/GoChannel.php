<?php
class GoChannel{

  protected $url;

  protected $postParams;

  public function __construct(){
    $config = Kiss_Registry::get("config");
    if($config->gochannel_url){
      $this->url = "http://".$config->gochannel_url;
    }else{
      throw new Exception("gochannel_url not config");
    }

    $this->postParams = array(
      "app_id"=>  $config->gochannel_app_id,
      "app_secret"=>  $config->gochannel_app_secret,
    );
  }

  public function getChannel($uid){
    //TODO set to Memcache
    $this->postParams["uid"] = $uid;
    $res = $this->curlPost($this->url."/api/get-channel",$this->postParams);
    return $res;
  }

  public function createChannel($uid){
    $this->postParams["uid"] = $uid;
    $res = $this->curlPost($this->url."/api/create-channel",$this->postParams);
    return $res;
  }

  public function push($uid,$msg){
    $this->postParams["uid"] = $uid;
    $this->postParams["message"] = $msg;
    $res = $this->curlPost($this->url."/api/push",$this->postParams);
    return $res;
  }

  public function broadcast($msg){
    $this->postParams["message"] = $msg;
    $res = $this->curlPost($this->url."/api/broadcast",$this->postParams);
    return $res;
  }

  public function close($uid){
    $this->postParams["uid"] = $uid;
    $res = $this->curlPost($this->url."/api/close-channel",$this->postParams);
    return $res;
  }

  protected function curlPost($url,$data){
    $ch = curl_init ();
    $postString = http_build_query ( $data, '&' );
    curl_setopt ( $ch, CURLOPT_URL, $url );
    curl_setopt ( $ch, CURLOPT_POST, 1 );
    curl_setopt ( $ch, CURLOPT_HEADER, 0 );
    curl_setopt ( $ch, CURLOPT_TIMEOUT, 5);
    curl_setopt ( $ch, CURLOPT_RETURNTRANSFER, 1 );
    curl_setopt ( $ch, CURLOPT_POSTFIELDS, $postString );
    Utils::log($postString."|||".$url,"curlPost");
    $return = curl_exec ( $ch );
    curl_close ( $ch );
    $result = json_decode($return,true);
    return $result?$result:false;
  }
}
