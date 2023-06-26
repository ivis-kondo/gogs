var dblClickFlag = null;

function ThroughDblClick() {
    // ダブルクリック（連続ポスト）の制御
    if (dblClickFlag == null) {
        dblClickFlag = 1;
        return true;
    } else {
        return false;
    }
}