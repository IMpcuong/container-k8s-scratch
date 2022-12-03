for d in $(ls -d */); do \
  cd $d; \
  for child in $(ls -d */); do cd $child; mkdir {0..100}; cd ..; done; \
  cd ..; \
done