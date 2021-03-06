.. _tarmak_clusters_plan:

tarmak clusters plan
--------------------

Plan changes on the currently configured cluster

Synopsis
~~~~~~~~


Plan changes on the currently configured cluster

::

  tarmak clusters plan [flags]

Options
~~~~~~~

::

  -h, --help                     help for plan
  -P, --plan-file-store string   location to store terraform plan executable file (default "${TARMAK_CONFIG}/${CURRENT_CLUSTER}/terraform/tarmak.plan")

Options inherited from parent commands
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

::

  -c, --config-directory string        config directory for tarmak's configuration (default "~/.tarmak")
      --current-cluster string         override the current cluster set in the config
      --keep-containers                do not clean-up terraform/packer containers after running them
      --log-flush-frequency duration   Maximum number of seconds between log flushes (default 5s)
      --public-api-endpoint            Override kubeconfig to point to cluster's public API endpoint
  -v, --verbose                        enable verbose logging
      --wing-dev-mode                  use a bundled wing version rather than a tagged release from GitHub

SEE ALSO
~~~~~~~~

* `tarmak clusters <tarmak_clusters.html>`_ 	 - Operations on clusters

