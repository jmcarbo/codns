 SUBDIRS = codns
      
.PHONY: subdirs $(SUBDIRS)
		      
subdirs: $(SUBDIRS)
	
$(SUBDIRS):
	$(MAKE) -C $@
																		     
