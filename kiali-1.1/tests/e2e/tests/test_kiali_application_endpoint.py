import tests.conftest as conftest

APPLICATION_TO_VALIDATE = 'productpage'
METRICS_PARAMS = {"direction": "outbound", "reporter": "destination"}

def test_application_list_endpoint(kiali_client):
    bookinfo_namespace = conftest.get_bookinfo_namespace()

    app_list = kiali_client.request(method_name='appList', path={'namespace': bookinfo_namespace, 'app': APPLICATION_TO_VALIDATE}).json()
    assert app_list != None
    for app in app_list.get('applications'):
      assert app.get('name') != None and app.get('name') != ''
      if 'traffic-generator' not in app.get('name'):
        assert app.get('istioSidecar') == True

    assert app_list.get('namespace').get('name') == bookinfo_namespace

def test_application_details_endpoint(kiali_client):
    bookinfo_namespace = conftest.get_bookinfo_namespace()

    app_details = kiali_client.request(method_name='appDetails', path={'namespace': bookinfo_namespace, 'app': APPLICATION_TO_VALIDATE}).json()

    assert app_details != None

    assert 'namespace' in app_details and app_details.get('namespace').get('name') == bookinfo_namespace

    assert 'workloads' in app_details
    workloads = app_details.get('workloads')
    assert len(workloads) > 0

    for workload in workloads:
      assert workload.get('istioSidecar') == True
      assert 'workloadName' in workload and len (workload.get('workloadName')) > 0


def __test_application_health_endpoint(kiali_client):
    bookinfo_namespace = conftest.get_bookinfo_namespace()

    app_health = kiali_client.request(method_name='appHealth', path={'namespace': bookinfo_namespace, 'app': APPLICATION_TO_VALIDATE}).json()
    assert app_health != None

    envoy = app_health.get('envoy')[0]
    assert envoy != None
    assert 'inbound' in envoy
    assert 'outbound' in envoy

    assert 'requests' in app_health
    assert 'workloadStatuses' in app_health

def test_application_metrics_endpoint(kiali_client):
    bookinfo_namespace = conftest.get_bookinfo_namespace()

    response = kiali_client.request(method_name='appMetrics', path={'namespace': bookinfo_namespace, 'app': APPLICATION_TO_VALIDATE}, params=METRICS_PARAMS)
    app_metrics = response.json()
    assert app_metrics != None

    metrics = app_metrics.get('metrics')
    assert 'request_count' in metrics
    assert 'request_error_count' in metrics
    assert 'tcp_received' in metrics
    assert 'tcp_sent' in metrics

    histograms = app_metrics.get('histograms')
    assert 'request_duration' in histograms
    assert 'request_size' in histograms
    assert 'response_size' in histograms
