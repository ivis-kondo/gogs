var dblClickFlag = null;

function PreventionDbClick() {
    if (dblClickFlag == null) {
        dblClickFlag = 1;
        console.log('PreventionDbClick() True');
        return true;
    } else {
        console.log('PreventionDbClick() Flase');
        return false;
    }
}