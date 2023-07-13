var dblClickFlag = null;

function PreventionDbClick() {
    if (dblClickFlag == null) {
        dblClickFlag = 1;
        return true;
    } else {
        return false;
    }
}