SHELL=/bin/bash

giorno=$(date +'%d')
m=$(date +'%m')
giornoDue=$((giorno+3))


case $m in
    1 | 3 | 5 | 7 | 8 | 10 | 12)
        togli=31
        ;;
    2) 
        togli=28
        ;;
    4 | 6 | 9 | 11) 
        togli=30
        ;;
esac

if [ ${giornoDue} > ${togli} ]
then
    giornoDue=$((giornoDue-togli))
    m=$((m+1))
    rm 30
fi
case $m in
    1)
        mese="gennaio"
        ;;
    2) 
        mese="febbraio"
        ;;
    3)
        mese="marzo"
            ;;
    4)
        mese="aprile"
        ;;
    5)
        mese="maggio"
            ;;
    6)
        mese="giugno"
            ;;
    7)
        mese="luglio"
            ;;
    8)
        mese="agosto"
            ;;
    9)
        mese="settembre"
            ;;
    10)
        mese="ottobre"
            ;;
    11)
        mese="novembre"
            ;;
    12)
        mese="dicembre"
            ;;
esac
echo "$giorno $giornoDue $mese"

mkdir "settimana 09 ($giorno-$giornoDue $mese)"